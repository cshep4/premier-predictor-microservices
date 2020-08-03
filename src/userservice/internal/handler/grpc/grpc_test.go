package grpc_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/handler/grpc"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/mocks/service"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"

	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	id1 = "1"
	id2 = "2"
)

func TestUserServiceServer_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	userService, err := grpc.New(service)
	require.NoError(t, err)

	t.Run("returns error if users cannot be retrieved", func(t *testing.T) {
		ctx := context.Background()

		service.EXPECT().GetAllUsers(ctx).Return(nil, errors.New("error"))

		resp, err := userService.GetAllUsers(ctx, &gen.GetAllUsersRequest{})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not get users", st.Message())
		assert.Nil(t, resp)
	})

	t.Run("returns all users", func(t *testing.T) {
		users := []*model.User{
			{
				Id: id1,
			},
			{
				Id: id2,
			},
		}
		ctx := context.Background()

		service.EXPECT().GetAllUsers(ctx).Return(users, nil)

		resp, err := userService.GetAllUsers(ctx, &gen.GetAllUsersRequest{})
		require.NoError(t, err)

		assert.Equal(t, id1, resp.Users[0].Id)
		assert.Equal(t, id2, resp.Users[1].Id)
	})
}

func TestUserServiceServer_GetAllUsersByIds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	userService, err := grpc.New(service)
	require.NoError(t, err)

	ids := []string{id1, id2}

	req := &gen.GetUsersByIdsRequest{
		Ids: ids,
	}

	t.Run("returns error if users cannot be retrieved", func(t *testing.T) {
		ctx := context.Background()

		service.EXPECT().GetAllUsersByIds(ctx, ids).Return(nil, errors.New("error"))

		resp, err := userService.GetUsersByIds(ctx, req)
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not get users", st.Message())
		assert.Nil(t, resp)
	})

	t.Run("returns all league users", func(t *testing.T) {
		users := []*model.User{
			{
				Id: id1,
			},
			{
				Id: id2,
			},
		}
		ctx := context.Background()

		service.EXPECT().GetAllUsersByIds(ctx, ids).Return(users, nil)

		resp, err := userService.GetUsersByIds(ctx, req)
		require.NoError(t, err)

		assert.Equal(t, id1, resp.Users[0].Id)
		assert.Equal(t, id2, resp.Users[1].Id)
	})
}

func TestUserServiceServer_GetOverallRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	userService, err := grpc.New(service)
	require.NoError(t, err)

	req := &gen.GetOverallRankRequest{
		Id: id1,
	}

	t.Run("returns error if rank cannot be retrieved", func(t *testing.T) {
		ctx := context.Background()

		service.EXPECT().GetOverallRank(ctx, id1).Return(int64(0), errors.New("error"))

		resp, err := userService.GetOverallRank(ctx, req)
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not get overall rank", st.Message())
		assert.Nil(t, resp)
	})

	t.Run("returns overall rank", func(t *testing.T) {
		rank := int64(1231)
		ctx := context.Background()
		service.EXPECT().GetOverallRank(ctx, id1).Return(rank, nil)

		resp, err := userService.GetOverallRank(ctx, req)
		require.NoError(t, err)

		assert.Equal(t, rank, resp.Rank)
	})
}

func TestUserServiceServer_GetRankForGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	userService, err := grpc.New(service)
	require.NoError(t, err)

	ids := []string{id1, id2}

	req := &gen.GetRankForGroupRequest{
		Id:  id1,
		Ids: ids,
	}

	t.Run("returns error if rank cannot be retrieved", func(t *testing.T) {
		ctx := context.Background()

		service.EXPECT().GetRankForGroup(ctx, id1, ids).Return(int64(0), errors.New("error"))

		resp, err := userService.GetRankForGroup(ctx, req)
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not get rank for group", st.Message())
		assert.Nil(t, resp)
	})

	t.Run("returns all league users", func(t *testing.T) {
		rank := int64(2)
		ctx := context.Background()

		service.EXPECT().GetRankForGroup(ctx, id1, ids).Return(rank, nil)

		resp, err := userService.GetRankForGroup(ctx, req)
		require.NoError(t, err)

		assert.Equal(t, rank, resp.Rank)
	})
}

func TestUserServiceServer_GetUserCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	userService, err := grpc.New(service)
	require.NoError(t, err)

	t.Run("returns error if user count cannot be retrieved", func(t *testing.T) {
		ctx := context.Background()

		service.EXPECT().GetUserCount(ctx).Return(int64(0), errors.New("error"))

		resp, err := userService.GetUserCount(ctx, &gen.GetUserCountRequest{})
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not get user count", st.Message())
		assert.Nil(t, resp)
	})

	t.Run("returns total user count", func(t *testing.T) {
		count := int64(2)
		ctx := context.Background()

		service.EXPECT().GetUserCount(ctx).Return(count, nil)

		resp, err := userService.GetUserCount(ctx, &gen.GetUserCountRequest{})
		require.NoError(t, err)

		assert.Equal(t, count, resp.Count)
	})
}

func TestUserServiceServer_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	userService, err := grpc.New(service)
	require.NoError(t, err)

	email := "📧"

	req := &gen.GetUserByEmailRequest{
		Email: email,
	}

	t.Run("returns error if user cannot be retrieved", func(t *testing.T) {
		ctx := context.Background()

		service.EXPECT().GetUserByEmail(ctx, email).Return(nil, errors.New("error"))

		resp, err := userService.GetUserByEmail(ctx, req)
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.Internal, st.Code())
		assert.Equal(t, "could not get user by email", st.Message())
		assert.Nil(t, resp)
	})

	t.Run("returns not found if user does not exist", func(t *testing.T) {
		ctx := context.Background()

		service.EXPECT().GetUserByEmail(ctx, email).Return(nil, fmt.Errorf("error: %w", model.ErrUserNotFound))

		resp, err := userService.GetUserByEmail(ctx, req)
		require.Error(t, err)

		st, ok := status.FromError(err)
		require.True(t, ok)

		assert.Equal(t, codes.NotFound, st.Code())
		assert.Equal(t, "user not found", st.Message())
		assert.Nil(t, resp)
	})

	t.Run("returns user", func(t *testing.T) {
		user := &model.User{
			Id: id1,
		}
		ctx := context.Background()

		service.EXPECT().GetUserByEmail(ctx, email).Return(user, nil)

		resp, err := userService.GetUserByEmail(ctx, req)
		require.NoError(t, err)

		assert.Equal(t, id1, resp.User.Id)
	})
}
