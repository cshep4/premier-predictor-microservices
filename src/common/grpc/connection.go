package grpc

import (
	"context"
	"github.com/cshep4/premier-predictor-microservices/src/common/grpc/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"log"
)

func Dial(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), options.ClientKeepAlive)
	if err != nil {
		return nil, err
	}
	onStateChange(conn, addr)

	return conn, nil
}

func onStateChange(conn *grpc.ClientConn, addr string) {
	go func() {
		for {
			conn.WaitForStateChange(context.Background(), conn.GetState())

			currentState := conn.GetState()
			log.Printf("%s connection state change - currentState: %s", addr, currentState)

			if currentState == connectivity.Connecting {
				continue
			}

			if currentState != connectivity.Ready {
				log.Printf("reconnecting %s connection", addr)

				var err error
				conn, err = grpc.Dial(addr, grpc.WithInsecure(), options.ClientKeepAlive)
				if err != nil {
					log.Printf("failed to reconnect %s connection", addr)
					continue
				}

				log.Printf("reconnected %s connection!", addr)
			}
		}
	}()
}
