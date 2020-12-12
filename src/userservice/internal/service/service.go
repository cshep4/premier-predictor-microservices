package user

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"

	"golang.org/x/crypto/bcrypt"
)

const (
	emailRegex = "^([_a-zA-Z0-9-]+(\\.[_a-zA-Z0-9-]+)*@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*(\\.[a-zA-Z]{1,6}))?$"
)

type (
	Store interface {
		GetUserById(ctx context.Context, id string) (*model.User, error)
		GetUserByEmail(ctx context.Context, email string) (*model.User, error)
		UpdateUserInfo(ctx context.Context, userInfo model.UserInfo) error
		UpdatePassword(ctx context.Context, id, password string) error
		UpdateSignature(ctx context.Context, id, signature string) error
		GetAllUsers(ctx context.Context) ([]*model.User, error)
		GetAllUsersByIds(ctx context.Context, ids []string) ([]*model.User, error)
		IsEmailTakenByADifferentUser(ctx context.Context, id, email string) bool
		GetOverallRank(ctx context.Context, id string) (int64, error)
		GetRankForGroup(ctx context.Context, id string, ids []string) (int64, error)
		GetUserCount(ctx context.Context) (int64, error)
		StoreUser(ctx context.Context, user model.User) (string, error)
	}

	service struct {
		store Store
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}

func New(store Store) (*service, error) {
	if store == nil {
		return nil, InvalidParameterError{Parameter: "store"}
	}

	return &service{
		store: store,
	}, nil
}

func (s *service) GetUserById(ctx context.Context, id string) (*model.User, error) {
	user, err := s.store.GetUserById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get_user_by_id: %w", err)
	}

	return user, nil
}

func (s *service) UpdateUserInfo(ctx context.Context, userInfo model.UserInfo) error {
	userInfo.Email = strings.ToLower(userInfo.Email)

	switch {
	case !regexp.MustCompile(emailRegex).MatchString(userInfo.Email):
		return model.InvalidParameterError{Parameter: "email"}
	case s.store.IsEmailTakenByADifferentUser(ctx, userInfo.Id, userInfo.Email):
		return model.InvalidParameterError{Parameter: "email already taken"}
	case userInfo.FirstName == "":
		return model.InvalidParameterError{Parameter: "first name"}
	case userInfo.Surname == "":
		return model.InvalidParameterError{Parameter: "surname"}
	}

	err := s.store.UpdateUserInfo(ctx, userInfo)
	if err != nil {
		return fmt.Errorf("update_user_info: %w", err)
	}

	return nil
}

func (s *service) UpdateUserPassword(ctx context.Context, updatePassword model.UpdatePassword) error {
	user, err := s.store.GetUserById(ctx, updatePassword.Id)
	if err != nil {
		return fmt.Errorf("get_user_by_id: %w", err)
	}

	switch {
	case bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(updatePassword.OldPassword)) != nil:
		return model.InvalidParameterError{Parameter: "old password does not match"}
	case updatePassword.NewPassword != updatePassword.ConfirmPassword:
		return model.InvalidParameterError{Parameter: "confirmation does not match"}
	case !validPassword(updatePassword.NewPassword):
		return model.InvalidParameterError{Parameter: "new password"}
	}

	newPassword, _ := bcrypt.GenerateFromPassword([]byte(updatePassword.NewPassword), 10)

	err = s.store.UpdatePassword(ctx, updatePassword.Id, string(newPassword))
	if err != nil {
		return fmt.Errorf("update_password: %w", err)
	}

	return nil
}

func validPassword(password string) bool {
	var sixToTwenty, number, upper, lower bool
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		}
	}

	sixToTwenty = len(password) >= 6 && len(password) <= 20

	return sixToTwenty && number && upper && lower
}

func (s *service) GetUserScore(ctx context.Context, id string) (int, error) {
	user, err := s.store.GetUserById(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("get_user_by_id: %w", err)
	}

	return user.Score, nil
}

func (s *service) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users, err := s.store.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("get_all_users: %w", err)
	}

	return users, nil
}

func (s *service) GetAllUsersByIds(ctx context.Context, ids []string) ([]*model.User, error) {
	users, err := s.store.GetAllUsersByIds(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("get_users_by_ids: %w", err)
	}

	return users, nil
}

func (s *service) GetRankForGroup(ctx context.Context, id string, ids []string) (int64, error) {
	rank, err := s.store.GetRankForGroup(ctx, id, ids)
	if err != nil {
		return 0, fmt.Errorf("get_rank_for_group: %w", err)
	}

	return rank, nil
}

func (s *service) GetOverallRank(ctx context.Context, id string) (int64, error) {
	rank, err := s.store.GetOverallRank(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("get_overall_rank: %w", err)
	}

	return rank, nil
}

func (s *service) GetUserCount(ctx context.Context) (int64, error) {
	count, err := s.store.GetUserCount(ctx)
	if err != nil {
		return 0, fmt.Errorf("get_user_count: %w", err)
	}

	return count, nil
}

func (s *service) StoreUser(ctx context.Context, user model.User) (string, error) {
	id, err := s.store.StoreUser(ctx, user)
	if err != nil {
		return "", fmt.Errorf("store_user: %w", err)
	}

	return id, nil
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := s.store.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("get_user_by_email: %w", err)
	}
	return user, nil
}

func (s *service) UpdatePassword(ctx context.Context, id, password string) error {
	err := s.store.UpdatePassword(ctx, id, password)
	if err != nil {
		return fmt.Errorf("update_password: %w", err)
	}

	return nil
}

func (s *service) UpdateSignature(ctx context.Context, id, signature string) error {
	err := s.store.UpdateSignature(ctx, id, signature)
	if err != nil {
		return fmt.Errorf("update_signature: %w", err)
	}

	return nil
}
