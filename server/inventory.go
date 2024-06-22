package server

import (
	"log"

	"github.com/natdanai0917/test_repo/modules/inventory/inventoryHandler"
	"github.com/natdanai0917/test_repo/modules/inventory/inventoryRepository"
	"github.com/natdanai0917/test_repo/modules/inventory/inventoryUsecase"
	"github.com/natdanai0917/test_repo/pkg/grpccon"

	inventoryPb "github.com/natdanai0917/test_repo/modules/inventory/inventoryPb"
)

func (s *server) inventoryService() {
	inventoryRepository := inventoryRepository.NewInventoryRepository(s.db)
	inventoryUsecase := inventoryUsecase.NewInventoryUsecase(inventoryRepository)
	inventoryHttpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, inventoryUsecase)
	inventoryGrpcHandler := inventoryHandler.NewInventoryGrpcHandler(inventoryHttpHandler)
	inventoryQueueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, inventoryHttpHandler)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.InventoryUrl)

		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, inventoryGrpcHandler)

		log.Printf("Inventory gRPC Server listening on %s", s.cfg.Grpc.InventoryUrl)
		grpcServer.Serve(lis)
	}()

	_ = inventoryHttpHandler
	_ = inventoryGrpcHandler
	_ = inventoryQueueHandler

	inventory := s.app.Group("/inventory_v1")

	//Health Check
	inventory.GET("", s.healthCheckService)
}
