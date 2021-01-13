package process

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/cshep4/data-structures/saga"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/event"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
)

type (
	payload struct {
		ID              string `json:"id"`
		FirstName       string `json:"firstName"`
		Surname         string `json:"surname"`
		Email           string `json:"email"`
		PredictedWinner string `json:"predictedWinner"`
	}
	publishEvent struct {
		topic           string
		publisher       event.Publisher
		user            model.User
		previousProcess string
	}
)

func NewPublishEvent(publisher event.Publisher, topic string, user model.User, previousProcess string) (*publishEvent, error) {
	switch {
	case publisher == nil:
		return nil, InvalidParameterError{Parameter: "publisher"}
	case topic == "":
		return nil, InvalidParameterError{Parameter: "topic"}
	case previousProcess == "":
		return nil, InvalidParameterError{Parameter: "topic"}
	}

	return &publishEvent{
		publisher:       publisher,
		topic:           topic,
		user:            user,
		previousProcess: previousProcess,
	}, nil
}

func (p publishEvent) Name() string {
	return "publish_event"
}

func (p publishEvent) Execute(ctx context.Context, results map[string]saga.Result) (saga.Result, error) {
	res, ok := results[p.previousProcess]
	if !ok {
		return nil, errors.New("cannot find previous process result")
	}

	id, ok := res.(string)
	if !ok {
		return nil, fmt.Errorf("invalid store_user result: %v", res)
	}

	b, err := json.Marshal(payload{
		ID:              id,
		FirstName:       p.user.FirstName,
		Surname:         p.user.Surname,
		Email:           p.user.Email,
		PredictedWinner: p.user.PredictedWinner,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot json marshal: %w", err)
	}

	_, err = p.publisher.PublishWithContext(ctx, &sns.PublishInput{
		Message:  aws.String(string(b)),
		TopicArn: aws.String(p.topic),
	})
	if err != nil {
		return nil, fmt.Errorf("cannot publish event: %w", err)
	}

	return nil, nil
}

func (p publishEvent) Rollback(context.Context, saga.ErrorHandler, map[string]saga.Result) error {
	return nil
}
