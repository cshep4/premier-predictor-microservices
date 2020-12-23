package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/app"
	"github.com/cshep4/premier-predictor-microservices/src/common/auth"
	grpcconn "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"github.com/cshep4/premier-predictor-microservices/src/common/runner/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/runner/http"
	"github.com/cshep4/premier-predictor-microservices/src/common/store/mongo"
	grpchandler "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/handler/grpc"
	httphandler "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/handler/http"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/prediction"
	svc "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/service"
	mongostore "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/store/mongo"
)

const (
	serviceName = "livematchservice"
	logLevel    = "info"
)

func start(ctx context.Context) error {
	var (
		authAddr       string
		predictionAddr string
		httpPortEnv    string
		grpcPortEnv    string
	)
	for k, v := range map[string]*string{
		"AUTH_ADDR":       &authAddr,
		"PREDICTION_ADDR": &predictionAddr,
		"HTTP_PORT":       &httpPortEnv,
		"PORT":            &grpcPortEnv,
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

	authConn, err := grpcconn.Dial(ctx, authAddr)
	if err != nil {
		return fmt.Errorf("create_auth_connection: %w", err)
	}
	defer authConn.Close()
	authClient := gen.NewAuthServiceClient(authConn)

	predictionConn, err := grpcconn.Dial(ctx, predictionAddr)
	if err != nil {
		return fmt.Errorf("create_prediction_connection: %w", err)
	}
	defer predictionConn.Close()
	predictionClient := gen.NewPredictionServiceClient(predictionConn)

	predictor, err := prediction.New(predictionClient)
	if err != nil {
		return fmt.Errorf("create_predictor: %w", err)
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

	service, err := svc.New(store, predictor)
	if err != nil {
		return fmt.Errorf("create_service: %w", err)
	}

	h, err := httphandler.New(service)
	if err != nil {
		return fmt.Errorf("create_http_handler: %w", err)
	}

	rpc, err := grpchandler.New(service, time.Minute)
	if err != nil {
		return fmt.Errorf("create_grpc_handler: %w", err)
	}

	authenticator, err := auth.New(authClient, "livematch", h)
	if err != nil {
		return fmt.Errorf("create_authenticator: %w", err)
	}

	//tracer := tracer.New()

	app := app.New(
		//app.WithStartupFunc(gcp.Profile(serviceName, version)),
		//app.WithStartupFunc(gcp.Trace),
		app.WithShutdownFunc(authConn.Close),
		app.WithShutdownFunc(predictionConn.Close),
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
		log.Error(ctx, "startup_error", log.ErrorParam(err))
	}
}
