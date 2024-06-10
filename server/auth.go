package server

import (
	"github.com/natdanai0917/test_repo/modules/auth/authHandler"
	"github.com/natdanai0917/test_repo/modules/auth/authRepository"
	"github.com/natdanai0917/test_repo/modules/auth/authUsecase"
)

func (s *server) authService() {
	authRepository := authRepository.NewAuthRepository(s.db)
	authUsecase := authUsecase.NewAuthUseCase(authRepository)
	authHttpHandler := authHandler.NewAuthHttpHandler(s.cfg, authUsecase)
	authGrpcHandler := authHandler.NewAuthGrpcHandler(authHttpHandler)

	_ = authHttpHandler
	_ = authGrpcHandler

	auth := s.app.Group("/auth_v1")

	//Health Check
	auth.GET("", s.healthCheckService)
}
