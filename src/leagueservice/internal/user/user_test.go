package user

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"strconv"
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/mocks/token"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/mocks/user"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"

	pb "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	id1   = "🆔"
	id2   = "⚽️"
	token = "🔑"
)

func TestUserService_GetAllUsers(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	userClient := user_mock.NewMockUserServiceClient(ctrl)
	tokenGenerator := token_mock.NewMockTokenGenerator(ctrl)

	service, err := New(tokenGenerator, userClient)
	require.NoError(t, err)

	var users []*pb.User
	var expectedUsers []model.LeagueUser

	for i := 0; i < 200000; i++ {
		users = append(users, &pb.User{
			Id: strconv.Itoa(i),
		})
		expectedUsers = append(expectedUsers, model.LeagueUser{
			Id: strconv.Itoa(i),
		})
	}

	resp := &pb.GetAllUsersResponse{
		Users: users,
	}

	t.Run("returns error if error generating token", func(t *testing.T) {
		tokenGenerator.EXPECT().Generate(ctx, "user").Return("", errors.New("error"))

		result, err := service.GetAllUsers(ctx)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem retrieving users", func(t *testing.T) {
		expectedCtx := metadata.AppendToOutgoingContext(ctx, "token", token)

		tokenGenerator.EXPECT().Generate(ctx, "user").Return(token, nil)
		userClient.EXPECT().GetAllUsers(expectedCtx, &pb.GetAllUsersRequest{}).Return(nil, errors.New("error"))

		result, err := service.GetAllUsers(ctx)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets all users from UserService", func(t *testing.T) {
		expectedCtx := metadata.AppendToOutgoingContext(ctx, "token", token)

		tokenGenerator.EXPECT().Generate(ctx, "user").Return(token, nil)
		userClient.EXPECT().GetAllUsers(expectedCtx, &pb.GetAllUsersRequest{}).Return(resp, nil)

		result, err := service.GetAllUsers(ctx)
		require.NoError(t, err)

		assert.Equal(t, expectedUsers, result)
	})
}

func TestUserService_GetLeagueUsers(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	userClient := user_mock.NewMockUserServiceClient(ctrl)
	tokenGenerator := token_mock.NewMockTokenGenerator(ctrl)

	service, err := New(tokenGenerator, userClient)
	require.NoError(t, err)

	var users []*pb.User
	var expectedUsers []model.LeagueUser

	for i := 0; i < 5; i++ {
		users = append(users, &pb.User{
			Id: strconv.Itoa(i),
		})
		expectedUsers = append(expectedUsers, model.LeagueUser{
			Id: strconv.Itoa(i),
		})
	}

	ids := []string{id1, id2}

	req := &pb.GetUsersByIdsRequest{
		Ids: ids,
	}

	resp := &pb.GetUsersByIdsResponse{
		Users: users,
	}

	t.Run("returns error if error generating token", func(t *testing.T) {
		tokenGenerator.EXPECT().Generate(ctx, "user").Return("", errors.New("error"))

		result, err := service.GetLeagueUsers(ctx, ids)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem retrieving users", func(t *testing.T) {
		expectedCtx := metadata.AppendToOutgoingContext(ctx, "token", token)

		tokenGenerator.EXPECT().Generate(ctx, "user").Return(token, nil)
		userClient.EXPECT().GetUsersByIds(expectedCtx, req).Return(nil, errors.New("error"))

		result, err := service.GetLeagueUsers(ctx, ids)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets all users from UserService", func(t *testing.T) {
		expectedCtx := metadata.AppendToOutgoingContext(ctx, "token", token)

		tokenGenerator.EXPECT().Generate(ctx, "user").Return(token, nil)
		userClient.EXPECT().GetUsersByIds(expectedCtx, req).Return(resp, nil)

		result, err := service.GetLeagueUsers(ctx, ids)
		require.NoError(t, err)

		assert.Equal(t, expectedUsers, result)
	})
}

func TestUserService_GetOverallRank(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	userClient := user_mock.NewMockUserServiceClient(ctrl)
	tokenGenerator := token_mock.NewMockTokenGenerator(ctrl)

	service, err := New(tokenGenerator, userClient)
	require.NoError(t, err)

	rank := int64(12345)

	req := &pb.GetOverallRankRequest{
		Id: id1,
	}

	resp := &pb.GetOverallRankResponse{
		Rank: rank,
	}

	t.Run("returns error if error generating token", func(t *testing.T) {
		tokenGenerator.EXPECT().Generate(ctx, "user").Return("", errors.New("error"))

		result, err := service.GetOverallRank(ctx, id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem retrieving users", func(t *testing.T) {
		expectedCtx := metadata.AppendToOutgoingContext(ctx, "token", token)

		tokenGenerator.EXPECT().Generate(ctx, "user").Return(token, nil)
		userClient.EXPECT().GetOverallRank(expectedCtx, req).Return(nil, errors.New("error"))

		result, err := service.GetOverallRank(ctx, id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets overall rank from UserService", func(t *testing.T) {
		expectedCtx := metadata.AppendToOutgoingContext(ctx, "token", token)

		tokenGenerator.EXPECT().Generate(ctx, "user").Return(token, nil)
		userClient.EXPECT().GetOverallRank(expectedCtx, req).Return(resp, nil)

		result, err := service.GetOverallRank(ctx, id1)
		require.NoError(t, err)

		assert.Equal(t, rank, result)
	})
}

func TestUserService_GetLeagueRank(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	userClient := user_mock.NewMockUserServiceClient(ctrl)
	tokenGenerator := token_mock.NewMockTokenGenerator(ctrl)

	service, err := New(tokenGenerator, userClient)
	require.NoError(t, err)

	rank := int64(12345)

	ids := []string{id1, id2}

	req := &pb.GetRankForGroupRequest{
		Id:  id1,
		Ids: ids,
	}

	resp := &pb.GetRankForGroupResponse{
		Rank: rank,
	}

	t.Run("returns error if error generating token", func(t *testing.T) {
		tokenGenerator.EXPECT().Generate(ctx, "user").Return("", errors.New("error"))

		result, err := service.GetLeagueRank(ctx, id1, ids)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem retrieving users", func(t *testing.T) {
		expectedCtx := metadata.AppendToOutgoingContext(ctx, "token", token)

		tokenGenerator.EXPECT().Generate(ctx, "user").Return(token, nil)
		userClient.EXPECT().GetRankForGroup(expectedCtx, req).Return(nil, errors.New("error"))

		result, err := service.GetLeagueRank(ctx, id1, ids)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets all users from UserService", func(t *testing.T) {
		expectedCtx := metadata.AppendToOutgoingContext(ctx, "token", token)

		tokenGenerator.EXPECT().Generate(ctx, "user").Return(token, nil)
		userClient.EXPECT().GetRankForGroup(expectedCtx, req).Return(resp, nil)

		result, err := service.GetLeagueRank(ctx, id1, ids)
		require.NoError(t, err)

		assert.Equal(t, rank, result)
	})
}

func TestUserService_GetUserCount(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	userClient := user_mock.NewMockUserServiceClient(ctrl)
	tokenGenerator := token_mock.NewMockTokenGenerator(ctrl)

	service, err := New(tokenGenerator, userClient)
	require.NoError(t, err)

	count := int64(1234)

	resp := &pb.GetUserCountResponse{
		Count: count,
	}

	t.Run("returns error if error generating token", func(t *testing.T) {
		tokenGenerator.EXPECT().Generate(ctx, "user").Return("", errors.New("error"))

		result, err := service.GetUserCount(ctx)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem retrieving users", func(t *testing.T) {
		expectedCtx := metadata.AppendToOutgoingContext(ctx, "token", token)

		tokenGenerator.EXPECT().Generate(ctx, "user").Return(token, nil)
		userClient.EXPECT().GetUserCount(expectedCtx, &pb.GetUserCountRequest{}).Return(nil, errors.New("error"))

		result, err := service.GetUserCount(ctx)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets all users from UserService", func(t *testing.T) {
		expectedCtx := metadata.AppendToOutgoingContext(ctx, "token", token)

		tokenGenerator.EXPECT().Generate(ctx, "user").Return(token, nil)
		userClient.EXPECT().GetUserCount(expectedCtx, &pb.GetUserCountRequest{}).Return(resp, nil)

		result, err := service.GetUserCount(ctx)
		require.NoError(t, err)

		assert.Equal(t, count, result)
	})
}
