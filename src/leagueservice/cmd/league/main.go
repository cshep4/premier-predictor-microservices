package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	pb "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/app"
	"github.com/cshep4/premier-predictor-microservices/src/common/auth"
	grpcconn "github.com/cshep4/premier-predictor-microservices/src/common/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"github.com/cshep4/premier-predictor-microservices/src/common/runner/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/common/runner/http"
	"github.com/cshep4/premier-predictor-microservices/src/common/store/mongo"

	grpcHandler "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/handler/grpc"
	httpHandler "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/handler/http"
	svc "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/service"
	leaguestore "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/store/league/mongo"
	overallstore "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/store/overall/mongo"
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

	client, err := mongo.New(ctx)
	if err != nil {
		return fmt.Errorf("create mongo client: %w", err)
	}

	leagueStore, err := leaguestore.New(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to create store: %w", err)
	}
	defer leagueStore.Close(ctx)

	userStore, err := overallstore.New(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to create store: %w", err)
	}
	defer leagueStore.Close(ctx)

	service, err := svc.New(leagueStore, userStore, timer{})
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
		//app.WithStartupFunc(gcp.Profile(serviceName, version)),
		//app.WithStartupFunc(gcp.Trace),
		app.WithShutdownFunc(authConn.Close),
		app.WithShutdownFuncContext(leagueStore.Close),
		app.WithShutdownFuncContext(userStore.Close),
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
