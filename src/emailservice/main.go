package main

import (
	"errors"
	"fmt"
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)

type emailService struct {
	sendGridApiKey string
}

func (e *emailService) Send(ctx context.Context, req *model.SendEmailRequest) (*empty.Empty, error) {
	from := mail.NewEmail(req.Sender, req.SenderEmail)
	to := mail.NewEmail(req.Recipient, req.RecipientEmail)

	subject := req.Subject

	plainTextContent := req.Content
	htmlContent := strings.Replace(req.Content, "\n", "<br>", -1)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(e.sendGridApiKey)

	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		err := errors.New(response.Body)
		log.Println(err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatalf("port not set")
	}

	sendGridApiKey, ok := os.LookupEnv("SENDGRID_API_KEY")
	if !ok {
		log.Fatalf("SendGrid API key not set")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	model.RegisterEmailServiceServer(grpcServer, &emailService{
		sendGridApiKey: sendGridApiKey,
	})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}
}
