package itemUsecase

import "github.com/natdanai0917/test_repo/modules/item/itemRepository"

type (
	ItemUsecaseService interface{}

	itemUsecase struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecase{itemRepository}
}
