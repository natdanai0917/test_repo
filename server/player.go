package server

import (
	"github.com/natdanai0917/test_repo/modules/player/playerHandler"
	"github.com/natdanai0917/test_repo/modules/player/playerRepository"
	"github.com/natdanai0917/test_repo/modules/player/playerUsecase"
)

func (s *server) playerService() {
	playerRepository := playerRepository.NewPlayerRepository(s.db)
	playerUsecase := playerUsecase.NewPlayerUsecase(playerRepository)
	playerHttpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, playerUsecase)
	playerGrpcHandler := playerHandler.NewPlayerGrpcHandler(playerHttpHandler)
	playerQueueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, playerHttpHandler)

	_ = playerHttpHandler
	_ = playerGrpcHandler
	_ = playerQueueHandler

	player := s.app.Group("/player_v1")

	//Health Check
	player.GET("/health", s.healthCheckService)
}
