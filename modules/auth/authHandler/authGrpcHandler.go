package authHandler

import (
	"context"

	authPb "github.com/natdanai0917/test_repo/modules/auth/authPb"
	"github.com/natdanai0917/test_repo/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandler {
	return &authGrpcHandler{authUsecase: authUsecase}
}

func (g *authGrpcHandler) CredentialSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return nil, nil
}

func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}
