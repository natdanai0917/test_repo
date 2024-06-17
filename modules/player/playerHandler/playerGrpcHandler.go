package playerHandler

import (
	"context"

	playerPb "github.com/natdanai0917/test_repo/modules/player/playerPb"
	"github.com/natdanai0917/test_repo/modules/player/playerUsecase"
)

type (
	playerGrpcHandler struct {
		playerPb      playerPb.UnimplementedPlayerGrpcServiceServer
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecaseService) *playerGrpcHandler {
	return &playerGrpcHandler{playerUsecase: playerUsecase}
}

func (g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func (g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func (g *playerGrpcHandler) GetPlayerSavingAccountReq(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	return nil, nil
}
