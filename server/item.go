package server

import (
	"github.com/natdanai0917/test_repo/modules/item/itemHandler"
	"github.com/natdanai0917/test_repo/modules/item/itemRepository"
	"github.com/natdanai0917/test_repo/modules/item/itemUsecase"
)

func (s *server) itemService() {
	itemRepository := itemRepository.NewItemRepository(s.db)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepository)
	itemHttpHandler := itemHandler.NewItemHttpHandler(s.cfg, itemUsecase)
	itemGrpcHandler := itemHandler.NewItemGrpcHandler(itemHttpHandler)

	_ = itemHttpHandler
	_ = itemGrpcHandler

	item := s.app.Group("/item_v1")

	//Health Check
	item.GET("", s.healthCheckService)
}
