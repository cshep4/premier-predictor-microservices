package process

import (
	"context"
	"fmt"

	"github.com/cshep4/data-structures/saga"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/store/mongo"
)

type updateUser struct {
	store            mongo.Store
	newUserInfo      model.UserInfo
	previousUserInfo model.UserInfo
}

func NewUpdateUser(store mongo.Store, newUserInfo, previousUserInfo model.UserInfo) (*updateUser, error) {
	if store == nil {
		return nil, InvalidParameterError{Parameter: "store"}
	}

	return &updateUser{
		newUserInfo:      newUserInfo,
		previousUserInfo: previousUserInfo,
		store:            store,
	}, nil
}

func (s updateUser) Name() string {
	return "update_user"
}

func (s updateUser) Execute(ctx context.Context, _ map[string]saga.Result) (saga.Result, error) {
	err := s.store.UpdateUserInfo(ctx, s.newUserInfo)
	if err != nil {
		return "", fmt.Errorf("cannot update user info: %w", err)
	}

	return s.newUserInfo.Id, nil
}

func (s updateUser) Rollback(ctx context.Context, _ saga.ErrorHandler, _ map[string]saga.Result) error {
	err := s.store.UpdateUserInfo(ctx, s.previousUserInfo)
	if err != nil {
		return fmt.Errorf("cannot revert user info: %w", err)
	}

	return nil
}
