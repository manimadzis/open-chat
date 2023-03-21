package channel_service

import (
	"context"
	"github.com/golang/mock/gomock"
	"open-chat/internal/entities"
	"open-chat/internal/repositories/mock"
	"open-chat/internal/services/role_checker"
	"testing"
)

func TestChannelService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	channelRepo := mock.NewMockChannelRepository(ctrl)
	ctx := context.Background()
	channel := &entities.Channel{CreatorId: 123, Server: &entities.Server{Id: 123, Name: "123"}, Name: "123"}
	channelRepo.EXPECT().Create(ctx, channel).Return(nil)
	roleChecker := role_checker.NewRoleChecker()
	NewChannelService(channelRepo)

}
