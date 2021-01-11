package process

import (
	"context"
	"errors"
	"fmt"

	"github.com/cshep4/data-structures/saga"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/store/mongo"
)

type (
	storeUser struct {
		store mongo.Store
		user  model.User
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}

func NewStoreUser(store mongo.Store, user model.User) (*storeUser, error) {
	if store == nil {
		return nil, InvalidParameterError{Parameter: "store"}
	}

	return &storeUser{
		user:  user,
		store: store,
	}, nil
}

func (s storeUser) Name() string {
	return "store_user"
}

func (s storeUser) Execute(ctx context.Context, _ map[string]saga.Result) (saga.Result, error) {
	id, err := s.store.StoreUser(ctx, s.user)
	if err != nil {
		return "", fmt.Errorf("cannot store user: %w", err)
	}

	return id, nil
}

func (s storeUser) Rollback(ctx context.Context, _ saga.ErrorHandler, results map[string]saga.Result) error {
	res, ok := results[s.Name()]
	if !ok {
		return errors.New("cannot find execute result")
	}

	id, ok := res.(string)
	if !ok {
		return fmt.Errorf("invalid execute result: %v", res)
	}

	err := s.store.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("cannot delete user: %w", err)
	}

	return nil
}
