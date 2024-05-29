package playerUsecase

import "github.com/natdanai0917/test_repo/modules/player/playerRepository"

type (
	PlayerUsecaseService interface {
	}

	playerUsecase struct {
		playerRepository playerRepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase(playerRepository playerRepository.PlayerRepositoryService) PlayerUsecaseService {
	return &playerUsecase{playerRepository}
}
