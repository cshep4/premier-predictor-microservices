package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cshep4/premier-predictor-microservices/src/common/gcp"
	"github.com/cshep4/premier-predictor-microservices/src/common/gcp/tracer"
	"github.com/cshep4/premier-predictor-microservices/src/common/store/mongo"
	"log"
	"os"
	"strconv"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/app"
	"github.com/cshep4/premier-predictor-microservices/src/common/auth"
	grpcconn "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/run"
	"github.com/cshep4/premier-predictor-microservices/src/common/runner/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/runner/http"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/fixture"
	grpchandler "github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/handler/grpc"
	httphandler "github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/handler/http"
	svc "github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/service"
	mongostore "github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/store/mongo"
	"golang.org/x/sync/errgroup"
)

const (
	serviceName = "predictionservice"
	version     = "1.0.0"
)

func start(ctx context.Context) error {
	var (
		authAddr    string
		fixtureAddr string
		httpPortEnv string
		grpcPortEnv string
	)
	for k, v := range map[string]*string{
		"AUTH_ADDR":    &authAddr,
		"FIXTURE_ADDR": &fixtureAddr,
		"HTTP_PORT":    &httpPortEnv,
		"PORT":         &grpcPortEnv,
	} {
		var ok bool
		if *v, ok = os.LookupEnv(k); !ok {
			return fmt.Errorf("missing_env_variable: %s", k)
		}
	}

	httpPort, err := strconv.Atoi(httpPortEnv)
	if err != nil {
		return errors.New("invalid_http_port")
	}

	grpcPort, err := strconv.Atoi(grpcPortEnv)
	if err != nil {
		return errors.New("invalid_grpc_port")
	}

	authConn, err := grpcconn.Dial(authAddr)
	if err != nil {
		return fmt.Errorf("create_auth_connection: %w", err)
	}
	defer authConn.Close()
	authClient := gen.NewAuthServiceClient(authConn)

	fixtureConn, err := grpcconn.Dial(fixtureAddr)
	if err != nil {
		return fmt.Errorf("create_fixture_connection: %w", err)
	}
	defer fixtureConn.Close()
	fixtureClient := gen.NewFixtureServiceClient(fixtureConn)

	authenticator, err := auth.New(authClient)
	if err != nil {
		return fmt.Errorf("create_authenticator: %w", err)
	}

	fixtureService, err := fixture.New(fixtureClient)
	if err != nil {
		return fmt.Errorf("create_fixture_client: %w", err)
	}

	client, err := mongo.New(ctx)
	if err != nil {
		return fmt.Errorf("create_mongo_client: %w", err)
	}

	store, err := mongostore.New(ctx, client)
	if err != nil {
		return fmt.Errorf("create_store: %w", err)
	}
	defer store.Close(ctx)

	service, err := svc.New(store, fixtureService)
	if err != nil {
		return fmt.Errorf("create_service: %w", err)
	}

	h, err := httphandler.New(service)
	if err != nil {
		return fmt.Errorf("create_http_handler: %w", err)
	}

	rpc, err := grpchandler.New(service)
	if err != nil {
		return fmt.Errorf("create_grpc_handler: %w", err)
	}

	tracer := tracer.New()

	app := app.New(
		app.WithStartupFunc(gcp.Profile(serviceName, version)),
		app.WithStartupFunc(gcp.Trace),
		app.WithShutdownFunc(authConn.Close),
		app.WithShutdownFunc(fixtureConn.Close),
		app.WithShutdownFuncContext(store.Close),
		app.WithRunner(
			grpc.New(
				grpc.WithPort(grpcPort),
				grpc.WithUnaryInterceptor(tracer.GrpcUnary),
				grpc.WithStreamInterceptor(tracer.GrpcStream),
				grpc.WithUnaryInterceptor(authenticator.GrpcUnaryInterceptor),
				grpc.WithStreamInterceptor(authenticator.GrpcStreamInterceptor),
				grpc.WithRegisterer(rpc),
			),
		),
		app.WithRunner(
			http.New(
				http.WithPort(httpPort),
				http.WithHandler(tracer),
				http.WithMiddleware(authenticator.HttpMiddleware),
				http.WithRouter(h),
				http.WithRegisterer(http.Health()),
			),
		),
	)

	ctx, cancel := context.WithCancel(ctx)
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error { return app.Run(ctx) })
	g.Go(run.HandleShutdown(g, ctx, cancel, app.Shutdown))

	return g.Wait()
}

func main() {
	if err := start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
