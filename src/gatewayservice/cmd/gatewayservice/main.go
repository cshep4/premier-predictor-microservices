package main

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/resolver"
	"github.com/cshep4/premier-predictor-microservices/src/gatewayservice/internal/resolver/mutation"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	var (
		port             string
		authAddress      string
		userAddress      string
		liveMatchAddress string
		chatAddress      string
	)

	for k, v := range map[string]*string{
		"PORT":            &port,
		"AUTH_ADDR":       &authAddress,
		"USER_ADDR":       &userAddress,
		"LIVE_MATCH_ADDR": &liveMatchAddress,
		"CHAT_ADDR":       &chatAddress,
	} {
		var ok bool
		if *v, ok = os.LookupEnv(k); !ok {
			log.Fatalf("missing environment variable %s\n", k)
		}
	}

	authConn := newGrpcConnection(authAddress)
	authServiceClient := gen.NewAuthServiceClient(authConn)
	userConn := newGrpcConnection(userAddress)
	userServiceClient := gen.NewUserServiceClient(userConn)
	liveMatchConn := newGrpcConnection(liveMatchAddress)
	liveMatchServiceClient := gen.NewLiveMatchServiceClient(liveMatchConn)
	chatConn := newGrpcConnection(chatAddress)
	chatServiceClient := gen.NewChatServiceClient(chatConn)

	r := resolver.NewRoot(
		resolver.RootQuery(
			resolver.QueryUserService(userServiceClient),
		),
		resolver.RootMutation(
			resolver.MutationAuth(
				mutation.NewAuth(
					mutation.AuthService(authServiceClient),
				),
			),
			resolver.MutationChat(
				mutation.NewChat(
					mutation.ChatService(chatServiceClient),
				),
			),
		),
		resolver.RootSubscription(
			resolver.SubscriptionAuthService(authServiceClient),
			resolver.SubscriptionLiveMatchService(liveMatchServiceClient),
		),
	)

	h, err := handler.New(
		handler.Resolver(r),
		handler.Middleware(),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error { return handler.Listen(ctx, ":"+port, h) })

	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		select {
		case <-c:
			cancel()
			return nil
		case <-ctx.Done():
			g.Go(authConn.Close)
			g.Go(userConn.Close)
			g.Go(liveMatchConn.Close)
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		log.Println(err)
	}
}

func newGrpcConnection(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
