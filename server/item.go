package server

import (
	"log"

	"github.com/natdanai0917/test_repo/modules/item/itemHandler"
	"github.com/natdanai0917/test_repo/modules/item/itemRepository"
	"github.com/natdanai0917/test_repo/modules/item/itemUsecase"
	"github.com/natdanai0917/test_repo/pkg/grpccon"

	itemPb "github.com/natdanai0917/test_repo/modules/item/itemPb"
)

func (s *server) itemService() {
	itemRepository := itemRepository.NewItemRepository(s.db)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepository)
	itemHttpHandler := itemHandler.NewItemHttpHandler(s.cfg, itemUsecase)
	itemGrpcHandler := itemHandler.NewItemGrpcHandler(itemHttpHandler)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, itemGrpcHandler)

		log.Printf("Item gRPC server listening on %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()

	_ = itemHttpHandler
	_ = itemGrpcHandler

	item := s.app.Group("/item_v1")

	//Health Check
	item.GET("", s.healthCheckService)
}
