package inventoryHandler

import (
	"context"

	inventoryPb "github.com/natdanai0917/test_repo/modules/inventory/inventoryPb"
	"github.com/natdanai0917/test_repo/modules/inventory/inventoryUsecase"
)

type (
	inventoryGrpcHandler struct {
		inventoryPb.UnimplementedInventoryGrpcServiceServer
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{inventoryUsecase: inventoryUsecase}
}

func (g *inventoryGrpcHandler) IsAvailableToSell(ctx context.Context, req *inventoryPb.IsAvailableToSellReq) (*inventoryPb.IsAvailableToSellRes, error) {
	return nil, nil
}
