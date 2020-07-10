package notification

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/util"
)

type notifier struct {
	client gen.NotificationServiceClient
}

func NewNotifier(client gen.NotificationServiceClient) (*notifier, error) {
	return &notifier{
		client: client,
	}, nil
}

func (n *notifier) Send(ctx context.Context, notification model.Notification, userIds ...string) error {
	not := &gen.Notification{
		Title:   notification.Title,
		Message: notification.Message,
	}

	metadata, err := util.CreateRequestMetadata(ctx)
	if err != nil {
		return err
	}

	switch len(userIds) {
	case 0:
		_, err = n.client.SendToAll(metadata, not)
	case 1:
		req := &gen.SingleRequest{
			UserId:       userIds[0],
			Notification: not,
		}
		_, err = n.client.Send(metadata, req)
	default:
		req := &gen.GroupRequest{
			UserIds:      userIds,
			Notification: not,
		}
		_, err = n.client.SendToGroup(metadata, req)
	}

	return err
}
