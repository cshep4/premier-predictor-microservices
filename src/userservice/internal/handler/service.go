package handler

import (
	"context"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
)

type Service interface {
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUserInfo(ctx context.Context, userDetails model.UserInfo) error
	UpdateUserPassword(ctx context.Context, updatePassword model.UpdatePassword) error
	GetUserScore(ctx context.Context, id string) (int, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetAllUsersByIds(ctx context.Context, ids []string) ([]*model.User, error)
	GetRankForGroup(ctx context.Context, id string, ids []string) (int64, error)
	GetOverallRank(ctx context.Context, id string) (int64, error)
	GetUserCount(ctx context.Context) (int64, error)
	StoreUser(ctx context.Context, user model.User) (string, error)
	UpdatePassword(ctx context.Context, id, password string) error
	UpdateSignature(ctx context.Context, id, signature string) error
}
