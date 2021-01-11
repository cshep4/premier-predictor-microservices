package event

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sns"
)

type Publisher interface {
	PublishWithContext(ctx context.Context, input *sns.PublishInput, opts ...request.Option) (*sns.PublishOutput, error)
}

func BuildTopic(region, accountID, topic string) string {
	return fmt.Sprintf("arn:aws:sns:%s:%s:%s", region, accountID, topic)
}
