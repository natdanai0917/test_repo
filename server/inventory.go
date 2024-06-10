package server

import (
	"github.com/natdanai0917/test_repo/modules/inventory/inventoryHandler"
	"github.com/natdanai0917/test_repo/modules/inventory/inventoryRepository"
	"github.com/natdanai0917/test_repo/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	inventoryRepository := inventoryRepository.NewInventoryRepository(s.db)
	inventoryUsecase := inventoryUsecase.NewInventoryUsecase(inventoryRepository)
	inventoryHttpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, inventoryUsecase)
	inventoryGrpcHandler := inventoryHandler.NewInventoryGrpcHandler(inventoryHttpHandler)
	inventoryQueueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, inventoryHttpHandler)

	_ = inventoryHttpHandler
	_ = inventoryGrpcHandler
	_ = inventoryQueueHandler

	inventory := s.app.Group("/inventory_v1")

	//Health Check
	inventory.GET("", s.healthCheckService)
}
