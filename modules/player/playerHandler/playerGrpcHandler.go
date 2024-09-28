package playerHandler

import (
	"context"

	playerPb "github.com/natdanai0917/test_repo/modules/player/playerPb"
	"github.com/natdanai0917/test_repo/modules/player/playerUsecase"
)

type (
	playerGrpcHandler struct {
		playerPb.UnimplementedPlayerGrpcServiceServer
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

// FindPlayerProfileToRefresh implements test_repo.PlayerGrpcServiceServer.
func (g *playerGrpcHandler) FindPlayerProfileToRefresh(context.Context, *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	panic("unimplemented")
}

// GetPlayerSavingAccount implements test_repo.PlayerGrpcServiceServer.
func (g *playerGrpcHandler) GetPlayerSavingAccount(context.Context, *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedPlayerGrpcServiceServer implements test_repo.PlayerGrpcServiceServer.
func (g *playerGrpcHandler) mustEmbedUnimplementedPlayerGrpcServiceServer() {
	panic("unimplemented")
}

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecaseService) *playerGrpcHandler {
	return &playerGrpcHandler{playerUsecase: playerUsecase}
}

func (g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	return g.playerUsecase.FindOnePlayerCredential(ctx,req.Password,req.Email)
}

func (g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	return g.playerUsecase.FindOnePlayerProfileToRefresh(ctx, req.PlayerId)
}

func (g *playerGrpcHandler) GetPlayerSavingAccountReq(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	return nil, nil
}
