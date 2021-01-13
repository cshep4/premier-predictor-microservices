package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/event"
	grpchandler "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/handler/grpc"
	httphandler "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/handler/http"
	usersaga "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/saga"
	svc "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/service"
	mongoStore "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/store/mongo"

	awsconfig "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/cshep4/data-structures/saga"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/app"
	"github.com/cshep4/premier-predictor-microservices/src/common/auth"
	grpcconn "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"github.com/cshep4/premier-predictor-microservices/src/common/runner/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/runner/http"
	"github.com/cshep4/premier-predictor-microservices/src/common/store/mongo"
)

const (
	serviceName = "userservice"
	logLevel    = "info"
)

func start(ctx context.Context) error {
	var (
		authAddr     string
		httpPortEnv  string
		grpcPortEnv  string
		awsRegion    string
		awsAccessKey string
		awsSecretKey string
		awsAccountID string
	)
	for k, v := range map[string]*string{
		"AUTH_ADDR":      &authAddr,
		"HTTP_PORT":      &httpPortEnv,
		"PORT":           &grpcPortEnv,
		"AWS_REGION":     &awsRegion,
		"AWS_ACCESS_KEY": &awsAccessKey,
		"AWS_SECRET_KEY": &awsSecretKey,
		"AWS_ACCOUNT_ID": &awsAccountID,
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

	runner, err := saga.New(
		saga.ErrHandler(usersaga.ErrorHandler{}),
		saga.RollbackHandler(usersaga.RollbackHandler{}),
	)
	if err != nil {
		return fmt.Errorf("failed to create saga runner: %w", err)
	}

	sess, err := session.NewSession(&awsconfig.Config{
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
		Region:      &awsRegion,
	})
	if err != nil {
		return fmt.Errorf("failed to create aws session: %w", err)
	}

	service, err := svc.New(
		store,
		runner,
		sns.New(sess),
		event.BuildTopic(awsRegion, awsAccountID, "UserCreated"),
		event.BuildTopic(awsRegion, awsAccountID, "UserUpdated"),
	)
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
		//app.WithStartupFunc(gcp.Profile(serviceName, version)),
		//app.WithStartupFunc(gcp.Trace),
		app.WithShutdownFunc(authConn.Close),
		app.WithShutdownFuncContext(store.Close),
		app.WithRunner(
			grpc.New(
				grpc.WithPort(grpcPort),
				grpc.WithLogger(serviceName, logLevel),
				//grpc.WithUnaryInterceptor(tracer.GrpcUnary),
				//grpc.WithStreamInterceptor(tracer.GrpcStream),
				grpc.WithUnaryInterceptor(authenticator.GrpcUnary),
				//grpc.WithStreamInterceptor(authenticator.GrpcStream),
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
