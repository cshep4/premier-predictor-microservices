package grpc

import (
	"context"

	"github.com/cshep4/premier-predictor-microservices/src/common/grpc/options"
	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

func Dial(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), options.ClientKeepAlive)
	if err != nil {
		return nil, err
	}
	onStateChange(ctx, conn, addr)

	return conn, nil
}

func onStateChange(ctx context.Context, conn *grpc.ClientConn, addr string) {
	go func() {
		for {
			conn.WaitForStateChange(context.Background(), conn.GetState())

			currentState := conn.GetState()
			log.Debug(ctx, "connection_state_change", log.SafeParam("currentState", currentState.String()))

			if currentState == connectivity.Connecting {
				continue
			}

			if currentState != connectivity.Ready {
				log.Debug(ctx, "reconnecting", log.SafeParam("connection", addr))

				var err error
				conn, err = grpc.Dial(addr, grpc.WithInsecure(), options.ClientKeepAlive)
				if err != nil {
					log.Error(ctx, "failed_to_reconnect", log.SafeParam("connection", addr))
					continue
				}

				log.Debug(ctx, "reconnected!", log.SafeParam("connection", addr))
			}
		}
	}()
}
