package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	grpcHandler "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/handler/grpc"
	httpHandler "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/handler/http"
	svc "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/service"
	mongoStore "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/store/mongo"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/table"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/token"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/user"

	pb "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
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
	serviceName = "leagueservice"
	version     = "1.0.0"
	logLevel    = "info"
)

func start(ctx context.Context) error {
	var (
		authAddr    string
		userAddr    string
		httpPortEnv string
		grpcPortEnv string
	)
	for k, v := range map[string]*string{
		"AUTH_ADDR": &authAddr,
		"USER_ADDR": &userAddr,
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
	authClient := pb.NewAuthServiceClient(authConn)

	userConn, err := grpcconn.Dial(ctx, userAddr)
	if err != nil {
		return fmt.Errorf("create user connection: %w", err)
	}
	defer userConn.Close()
	userClient := pb.NewUserServiceClient(userConn)

	client, err := mongo.New(ctx)
	if err != nil {
		return fmt.Errorf("create mongo client: %w", err)
	}

	store, err := mongoStore.New(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to create store: %w", err)
	}
	defer store.Close(ctx)

	tokenGenerator, err := token.New(authClient)
	if err != nil {
		return fmt.Errorf("failed to create token generator: %w", err)
	}

	userService, err := user.New(tokenGenerator, userClient)
	if err != nil {
		return fmt.Errorf("failed to create userService: %w", err)
	}

	overallTable, err := table.NewOverallTable(ctx, userService)
	if err != nil {
		return fmt.Errorf("failed to create overallTable: %w", err)
	}

	service, err := svc.New(store, userService, overallTable, timer{})
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	h, err := httpHandler.New(service)
	if err != nil {
		return fmt.Errorf("create http handler: %w", err)
	}

	rpc, err := grpcHandler.New(service)
	if err != nil {
		return fmt.Errorf("create grpc handler: %w", err)
	}

	authenticator, err := auth.New(authClient, "league", h)
	if err != nil {
		return fmt.Errorf("create authenticator: %w", err)
	}

	//tracer := tracer.New()

	app := app.New(
		app.WithStartupFunc(gcp.Profile(serviceName, version)),
		app.WithStartupFunc(gcp.Trace),
		app.WithShutdownFunc(authConn.Close),
		app.WithShutdownFunc(userConn.Close),
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

type timer struct{}

func (timer) Now() time.Time {
	return time.Now()
}
