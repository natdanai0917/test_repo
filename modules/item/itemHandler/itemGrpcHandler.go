package itemHandler

import (
	"context"

	itemPb "github.com/natdanai0917/test_repo/modules/item/itemPb"
	"github.com/natdanai0917/test_repo/modules/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemPb.UnimplementedItemGrpcServiceServer
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{itemUsecase: itemUsecase}
}

func (g *itemGrpcHandler) FindItemInIds(ctx context.Context, req *itemPb.FindItemInIdsReq) (*itemPb.FindItemInIdsRes, error) {
	return nil, nil
}
