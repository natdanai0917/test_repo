package inventory

import (
	"github.com/natdanai0917/test_repo/modules/item"
	"github.com/natdanai0917/test_repo/modules/models"
)

type (
	UpdateInventoryReq struct {
		PlayerId string `json:"playerId" validate:"required,max=64"`
		ItemId   string `json:"item_id" validate:"required,max=64"`
	}

	ItemInventory struct {
		InventoryId string `json:"inventory_id"`
		*item.ItemShowCase
	}

	PlayerInventory struct {
		PlayerId string `json:"player_id"`
		*models.PaginateRes
	}
)
