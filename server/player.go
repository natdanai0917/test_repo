package server

import (
	"log"

	"github.com/natdanai0917/test_repo/modules/player/playerHandler"
	"github.com/natdanai0917/test_repo/modules/player/playerRepository"
	"github.com/natdanai0917/test_repo/modules/player/playerUsecase"
	"github.com/natdanai0917/test_repo/pkg/grpccon"

	playerPb "github.com/natdanai0917/test_repo/modules/player/playerPb"
)

func (s *server) playerService() {
	playerRepository := playerRepository.NewPlayerRepository(s.db)
	playerUsecase := playerUsecase.NewPlayerUsecase(playerRepository)
	playerHttpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, playerUsecase)
	playerGrpcHandler := playerHandler.NewPlayerGrpcHandler(playerUsecase)
	playerQueueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, playerUsecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)

		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, playerGrpcHandler)

		log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()

	_ = playerGrpcHandler
	_ = playerQueueHandler

	player := s.app.Group("/player_v1")

	//Health Check
	player.GET("", s.healthCheckService)

	player.POST("/player/register", playerHttpHandler.CreatePlayer)
	player.GET("/player/:player_id", playerHttpHandler.FineOnePlayerProfile)
	player.POST("/player/add-money", playerHttpHandler.AddPlayerMoney)
}
