package authHandler

import "github.com/natdanai0917/test_repo/modules/auth/authUsecase"

type (
	authGrpcHandler struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandler {
	return &authGrpcHandler{authUsecase}
}
