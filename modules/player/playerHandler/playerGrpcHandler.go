package playerHandler

import (
	"github.com/natdanai0917/test_repo/modules/player/playerUsecase"
)

type (
	playerGrpcHandler struct {
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecaseService) *playerGrpcHandler {
	return &playerGrpcHandler{playerUsecase}
}
