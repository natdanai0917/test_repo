package inventoryHandler

import (
	"github.com/natdanai0917/test_repo/config"
	"github.com/natdanai0917/test_repo/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface {
	}

	inventoryQueueHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryQueueHandlerService {
	return &inventoryQueueHandler{cfg, inventoryUsecase}
}
