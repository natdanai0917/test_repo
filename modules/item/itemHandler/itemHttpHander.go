package itemHandler

import (
	"github.com/natdanai0917/test_repo/config"
	"github.com/natdanai0917/test_repo/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface {
	}

	itemHttpHandler struct {
		cfg         *config.Config
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemUsecase.ItemUsecaseService) ItemHttpHandlerService {
	return &itemHttpHandler{cfg, itemUsecase}
}
