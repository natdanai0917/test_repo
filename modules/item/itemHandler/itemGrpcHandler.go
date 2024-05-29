package itemHandler

import "github.com/natdanai0917/test_repo/modules/item/itemUsecase"

type (
	itemGrpcHandler struct {
		authUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{itemUsecase}
}
