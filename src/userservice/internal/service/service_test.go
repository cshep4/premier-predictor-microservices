package user

import (
	"context"
	"errors"
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/mocks/store"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

const (
	userId           = "1"
	emailAddress     = "example@test.com"
	oldPassword      = "old password"
	newPassword      = "new password"
	newValidPassword = "Qwerty123"
)

var (
	testErr = errors.New("error")
)

func TestService_GetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := store_mock.NewMockStore(ctrl)

	service, err := New(store)
	require.NoError(t, err)

	t.Run("returns error if error getting user", func(t *testing.T) {
		store.EXPECT().GetUserById(context.Background(), userId).Return(nil, testErr)

		result, err := service.GetUserById(context.Background(), userId)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
		assert.Nil(t, result)
	})

	t.Run("gets user from db", func(t *testing.T) {
		user := &model.User{
			Id: userId,
		}

		store.EXPECT().GetUserById(context.Background(), userId).Return(user, nil)

		result, err := service.GetUserById(context.Background(), userId)

		require.NoError(t, err)
		assert.Equal(t, user, result)
	})
}

func TestService_UpdateUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := store_mock.NewMockStore(ctrl)

	service, err := New(store)
	require.NoError(t, err)

	t.Run("returns error if the email address is not valid", func(t *testing.T) {
		userInfo := model.UserInfo{
			Email: "invalid email address",
		}

		err := service.UpdateUserInfo(context.Background(), userInfo)
		require.Error(t, err)

		ipe, ok := err.(model.InvalidParameterError)
		require.True(t, ok)

		assert.Equal(t, "email", ipe.Parameter)
	})

	t.Run("returns error if the email address is already taken by a different user", func(t *testing.T) {
		userInfo := model.UserInfo{
			Id:    userId,
			Email: emailAddress,
		}

		store.EXPECT().IsEmailTakenByADifferentUser(context.Background(), userId, emailAddress).Return(true)

		err := service.UpdateUserInfo(context.Background(), userInfo)
		require.Error(t, err)

		ipe, ok := err.(model.InvalidParameterError)
		require.True(t, ok)

		assert.Equal(t, "email already taken", ipe.Parameter)
	})

	t.Run("returns error if the first name is blank", func(t *testing.T) {
		userInfo := model.UserInfo{
			Id:    userId,
			Email: emailAddress,
		}

		store.EXPECT().IsEmailTakenByADifferentUser(context.Background(), userId, emailAddress).Return(false)

		err := service.UpdateUserInfo(context.Background(), userInfo)
		require.Error(t, err)

		ipe, ok := err.(model.InvalidParameterError)
		require.True(t, ok)

		assert.Equal(t, "first name", ipe.Parameter)
	})

	t.Run("returns error if the surname is blank", func(t *testing.T) {
		userInfo := model.UserInfo{
			Id:        userId,
			Email:     emailAddress,
			FirstName: "first name",
		}

		store.EXPECT().IsEmailTakenByADifferentUser(context.Background(), userId, emailAddress).Return(false)

		err := service.UpdateUserInfo(context.Background(), userInfo)
		require.Error(t, err)

		ipe, ok := err.(model.InvalidParameterError)
		require.True(t, ok)

		assert.Equal(t, "surname", ipe.Parameter)
	})

	t.Run("returns error if details cannot be updated", func(t *testing.T) {
		userInfo := model.UserInfo{
			Id:        userId,
			Email:     emailAddress,
			FirstName: "first name",
			Surname:   "surname",
		}

		store.EXPECT().IsEmailTakenByADifferentUser(context.Background(), userId, emailAddress).Return(false)
		store.EXPECT().UpdateUserInfo(context.Background(), userInfo).Return(testErr)

		err := service.UpdateUserInfo(context.Background(), userInfo)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
	})

	t.Run("returns nil if details are updated successfully", func(t *testing.T) {
		userInfo := model.UserInfo{
			Id:        userId,
			Email:     emailAddress,
			FirstName: "first name",
			Surname:   "surname",
		}

		store.EXPECT().IsEmailTakenByADifferentUser(context.Background(), userId, emailAddress).Return(false)
		store.EXPECT().UpdateUserInfo(context.Background(), userInfo).Return(nil)

		err := service.UpdateUserInfo(context.Background(), userInfo)
		require.NoError(t, err)

		assert.Nil(t, err)
	})
}

func TestService_UpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := store_mock.NewMockStore(ctrl)

	service, err := New(store)
	require.NoError(t, err)

	t.Run("returns error if the user cannot be retrieved", func(t *testing.T) {
		store.EXPECT().GetUserById(context.Background(), userId).Return(nil, testErr)

		err := service.UpdateUserPassword(context.Background(), model.UpdatePassword{Id: userId})
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
	})

	t.Run("returns error if old password does not match what is currently stored", func(t *testing.T) {
		user := &model.User{
			Id:       userId,
			Password: hashPassword(oldPassword),
		}

		store.EXPECT().GetUserById(context.Background(), userId).Return(user, nil)

		updatePassword := model.UpdatePassword{
			Id:          userId,
			OldPassword: "different old password",
		}

		err := service.UpdateUserPassword(context.Background(), updatePassword)
		require.Error(t, err)

		ipe, ok := err.(model.InvalidParameterError)
		require.True(t, ok)

		assert.Equal(t, "old password does not match", ipe.Parameter)
	})

	t.Run("returns error if new password does not match the confirmation", func(t *testing.T) {
		user := &model.User{
			Id:       userId,
			Password: hashPassword(oldPassword),
		}

		store.EXPECT().GetUserById(context.Background(), userId).Return(user, nil)

		updatePassword := model.UpdatePassword{
			Id:              userId,
			OldPassword:     oldPassword,
			NewPassword:     newPassword,
			ConfirmPassword: "different confirmation password",
		}

		err := service.UpdateUserPassword(context.Background(), updatePassword)
		require.Error(t, err)

		ipe, ok := err.(model.InvalidParameterError)
		require.True(t, ok)

		assert.Equal(t, "confirmation does not match", ipe.Parameter)
	})

	t.Run("returns error if new password is not valid", func(t *testing.T) {
		user := &model.User{
			Id:       userId,
			Password: hashPassword(oldPassword),
		}

		store.EXPECT().GetUserById(context.Background(), userId).Return(user, nil)

		updatePassword := model.UpdatePassword{
			Id:              userId,
			OldPassword:     oldPassword,
			NewPassword:     newPassword,
			ConfirmPassword: newPassword,
		}

		err := service.UpdateUserPassword(context.Background(), updatePassword)
		require.Error(t, err)

		ipe, ok := err.(model.InvalidParameterError)
		require.True(t, ok)

		assert.Equal(t, "new password", ipe.Parameter)
	})

	t.Run("returns error if password cannot be updated", func(t *testing.T) {
		user := &model.User{
			Id:       userId,
			Password: hashPassword(oldPassword),
		}

		store.EXPECT().GetUserById(context.Background(), userId).Return(user, nil)

		updatePassword := model.UpdatePassword{
			Id:              userId,
			OldPassword:     oldPassword,
			NewPassword:     newValidPassword,
			ConfirmPassword: newValidPassword,
		}

		store.EXPECT().UpdatePassword(context.Background(), userId, gomock.Any()).Return(testErr)

		err := service.UpdateUserPassword(context.Background(), updatePassword)
		require.Error(t, err)

		assert.True(t, errors.Is(err, testErr))
	})

	t.Run("returns nil if password is updated successfully", func(t *testing.T) {
		user := &model.User{
			Id:       userId,
			Password: hashPassword(oldPassword),
		}

		store.EXPECT().GetUserById(context.Background(), userId).Return(user, nil)

		updatePassword := model.UpdatePassword{
			Id:              userId,
			OldPassword:     oldPassword,
			NewPassword:     newValidPassword,
			ConfirmPassword: newValidPassword,
		}

		store.EXPECT().UpdatePassword(context.Background(), userId, gomock.Any()).Return(nil)

		err := service.UpdateUserPassword(context.Background(), updatePassword)
		require.NoError(t, err)

		assert.Nil(t, err)
	})
}

func TestService_GetUserScore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := store_mock.NewMockStore(ctrl)

	service, err := New(store)
	require.NoError(t, err)

	t.Run("returns error if user cannot be retrieved", func(t *testing.T) {
		store.EXPECT().GetUserById(context.Background(), userId).Return(nil, testErr)

		result, err := service.GetUserScore(context.Background(), userId)

		require.Error(t, err)
		assert.True(t, errors.Is(err, testErr))
		assert.Equal(t, 0, result)
	})

	t.Run("returns specified user's score", func(t *testing.T) {
		const score = 1234

		user := &model.User{
			Id:    userId,
			Score: score,
		}

		store.EXPECT().GetUserById(context.Background(), userId).Return(user, nil)

		result, err := service.GetUserScore(context.Background(), userId)

		require.NoError(t, err)
		assert.Equal(t, score, result)
	})
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func TestService_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := store_mock.NewMockStore(ctrl)

	service, err := New(store)
	require.NoError(t, err)

	t.Run("returns error if users cannot be retrieved", func(t *testing.T) {
		store.EXPECT().GetAllUsers(context.Background()).Return(nil, testErr)

		result, err := service.GetAllUsers(context.Background())

		require.Error(t, err)
		assert.True(t, errors.Is(err, testErr))
		assert.Nil(t, result)
	})

	t.Run("returns all users", func(t *testing.T) {
		users := []*model.User{
			{
				Id: userId,
			},
		}

		store.EXPECT().GetAllUsers(context.Background()).Return(users, nil)

		result, err := service.GetAllUsers(context.Background())

		require.NoError(t, err)

		assert.Equal(t, users, result)
	})
}

func TestService_GetAllUsersByIds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := store_mock.NewMockStore(ctrl)

	service, err := New(store)
	require.NoError(t, err)

	ids := []string{userId}

	t.Run("returns error if users cannot be retrieved", func(t *testing.T) {
		store.EXPECT().GetAllUsersByIds(context.Background(), ids).Return(nil, testErr)

		result, err := service.GetAllUsersByIds(context.Background(), ids)

		require.Error(t, err)
		assert.True(t, errors.Is(err, testErr))
		assert.Nil(t, result)
	})

	t.Run("returns all users", func(t *testing.T) {
		users := []*model.User{
			{
				Id: userId,
			},
		}

		store.EXPECT().GetAllUsersByIds(context.Background(), ids).Return(users, nil)

		result, err := service.GetAllUsersByIds(context.Background(), ids)

		require.NoError(t, err)

		assert.Equal(t, users, result)
	})
}

func TestService_GetRankForGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := store_mock.NewMockStore(ctrl)

	service, err := New(store)
	require.NoError(t, err)

	ids := []string{userId}

	t.Run("returns error if error getting rank", func(t *testing.T) {
		store.EXPECT().GetRankForGroup(context.Background(), userId, ids).Return(int64(0), testErr)

		result, err := service.GetRankForGroup(context.Background(), userId, ids)

		require.Error(t, err)
		assert.True(t, errors.Is(err, testErr))
		assert.Empty(t, result)
	})

	t.Run("gets rank for group", func(t *testing.T) {
		store.EXPECT().GetRankForGroup(context.Background(), userId, ids).Return(int64(1), nil)

		result, err := service.GetRankForGroup(context.Background(), userId, ids)

		require.NoError(t, err)
		assert.Equal(t, int64(1), result)
	})
}

func TestService_GetOverallRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := store_mock.NewMockStore(ctrl)

	service, err := New(store)
	require.NoError(t, err)

	t.Run("returns error if there is a problem", func(t *testing.T) {
		store.EXPECT().GetOverallRank(context.Background(), userId).Return(int64(0), testErr)

		result, err := service.GetOverallRank(context.Background(), userId)

		require.Error(t, err)
		assert.True(t, errors.Is(err, testErr))
		assert.Empty(t, result)
	})

	t.Run("gets overall rank", func(t *testing.T) {
		store.EXPECT().GetOverallRank(context.Background(), userId).Return(int64(1), nil)

		result, err := service.GetOverallRank(context.Background(), userId)

		require.NoError(t, err)
		assert.Equal(t, int64(1), result)
	})
}

func TestService_GetUserCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := store_mock.NewMockStore(ctrl)

	service, err := New(store)
	require.NoError(t, err)

	t.Run("returns error if there is a problem", func(t *testing.T) {
		store.EXPECT().GetUserCount(context.Background()).Return(int64(0), testErr)

		result, err := service.GetUserCount(context.Background())

		require.Error(t, err)
		assert.True(t, errors.Is(err, testErr))
		assert.Empty(t, result)
	})

	t.Run("gets user count", func(t *testing.T) {
		store.EXPECT().GetUserCount(context.Background()).Return(int64(1), nil)

		result, err := service.GetUserCount(context.Background())

		require.NoError(t, err)
		assert.Equal(t, int64(1), result)
	})
}

func TestService_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := store_mock.NewMockStore(ctrl)

	service, err := New(store)
	require.NoError(t, err)

	t.Run("returns error if there is a problem", func(t *testing.T) {
		store.EXPECT().GetUserByEmail(context.Background(), emailAddress).Return(nil, testErr)

		result, err := service.GetUserByEmail(context.Background(), emailAddress)

		require.Error(t, err)
		assert.True(t, errors.Is(err, testErr))
		assert.Nil(t, result)
	})

	t.Run("gets user from db", func(t *testing.T) {
		user := &model.User{
			Id:    userId,
			Email: emailAddress,
		}

		store.EXPECT().GetUserByEmail(context.Background(), emailAddress).Return(user, nil)

		result, err := service.GetUserByEmail(context.Background(), emailAddress)

		require.NoError(t, err)
		assert.Equal(t, user, result)
	})
}
