package playerHandler

import (
	"github.com/natdanai0917/test_repo/config"
	"github.com/natdanai0917/test_repo/modules/player/playerUsecase"
)

type (
	PlayerHttpHandlerService interface {
	}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerHttpHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{cfg, playerUsecase}
}
