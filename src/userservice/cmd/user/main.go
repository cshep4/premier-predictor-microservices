package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	grpchandler "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/handler/grpc"
	httphandler "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/handler/http"
	svc "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/service"
	mongoStore "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/store/mongo"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/app"
	"github.com/cshep4/premier-predictor-microservices/src/common/auth"
	"github.com/cshep4/premier-predictor-microservices/src/common/gcp"
	grpcconn "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"github.com/cshep4/premier-predictor-microservices/src/common/runner/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/runner/http"
	"github.com/cshep4/premier-predictor-microservices/src/common/store/mongo"
)

const (
	serviceName = "userservice"
	version     = "1.0.0"
	logLevel    = "info"
)

func start(ctx context.Context) error {
	var (
		authAddr    string
		httpPortEnv string
		grpcPortEnv string
	)
	for k, v := range map[string]*string{
		"AUTH_ADDR": &authAddr,
		"HTTP_PORT": &httpPortEnv,
		"PORT":      &grpcPortEnv,
	} {
		var ok bool
		if *v, ok = os.LookupEnv(k); !ok {
			return fmt.Errorf("missing env variable: %s", k)
		}
	}

	httpPort, err := strconv.Atoi(httpPortEnv)
	if err != nil {
		return errors.New("invalid http port")
	}

	grpcPort, err := strconv.Atoi(grpcPortEnv)
	if err != nil {
		return errors.New("invalid grpc port")
	}

	authConn, err := grpcconn.Dial(ctx, authAddr)
	if err != nil {
		return fmt.Errorf("create auth connection: %w", err)
	}
	defer authConn.Close()
	authClient := gen.NewAuthServiceClient(authConn)

	client, err := mongo.New(ctx)
	if err != nil {
		return fmt.Errorf("create mongo client: %w", err)
	}

	store, err := mongoStore.New(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to create repository: %w", err)
	}
	defer store.Close(ctx)

	service, err := svc.New(store)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	h, err := httphandler.New(service)
	if err != nil {
		return fmt.Errorf("create http handler: %w", err)
	}

	rpc, err := grpchandler.New(service)
	if err != nil {
		return fmt.Errorf("create grpc handler: %w", err)
	}

	authenticator, err := auth.New(authClient, "user", h)
	if err != nil {
		return fmt.Errorf("create authenticator: %w", err)
	}

	//tracer := tracer.New()

	app := app.New(
		app.WithStartupFunc(gcp.Profile(serviceName, version)),
		app.WithStartupFunc(gcp.Trace),
		app.WithShutdownFunc(authConn.Close),
		app.WithShutdownFuncContext(store.Close),
		app.WithRunner(
			grpc.New(
				grpc.WithPort(grpcPort),
				grpc.WithLogger(serviceName, logLevel),
				//grpc.WithUnaryInterceptor(tracer.GrpcUnary),
				//grpc.WithStreamInterceptor(tracer.GrpcStream),
				grpc.WithUnaryInterceptor(authenticator.GrpcUnary),
				grpc.WithStreamInterceptor(authenticator.GrpcStream),
				grpc.WithRegisterer(rpc),
			),
		),
		app.WithRunner(
			http.New(
				http.WithPort(httpPort),
				http.WithLogger(serviceName, logLevel),
				//http.WithHandler(tracer),
				http.WithMiddleware(authenticator.Http),
				http.WithRouter(h),
				http.WithRegisterer(http.Health()),
			),
		),
	)

	return app.Run(ctx)
}

func main() {
	ctx := log.WithServiceName(context.Background(), log.New(logLevel), serviceName)
	if err := start(ctx); err != nil {
		log.Error(ctx, "error_starting_server", log.ErrorParam(err))
	}
}
