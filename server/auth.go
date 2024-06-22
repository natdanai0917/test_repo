package server

import (
	"log"

	"github.com/natdanai0917/test_repo/modules/auth/authHandler"
	"github.com/natdanai0917/test_repo/modules/auth/authRepository"
	"github.com/natdanai0917/test_repo/modules/auth/authUsecase"
	"github.com/natdanai0917/test_repo/pkg/grpccon"

	authPb "github.com/natdanai0917/test_repo/modules/auth/authPb"
)

func (s *server) authService() {
	authRepository := authRepository.NewAuthRepository(s.db)
	authUsecase := authUsecase.NewAuthUseCase(authRepository)
	authHttpHandler := authHandler.NewAuthHttpHandler(s.cfg, authUsecase)
	authGrpcHandler := authHandler.NewAuthGrpcHandler(authHttpHandler)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, authGrpcHandler)

		log.Printf("Auth gRPC Server listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = authHttpHandler
	_ = authGrpcHandler

	auth := s.app.Group("/auth_v1")

	//Health Check
	auth.GET("", s.healthCheckService)
}
