package middlewareHandler

import (
	"github.com/natdanai0917/test_repo/config"
	"github.com/natdanai0917/test_repo/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareUsecaseService interface{}

	middlewareHandler struct {
		cfg                      *config.Config
		middlewareUsecaseService middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareUsecaseService {
	return &middlewareHandler{cfg, middlewareUsecase}
}
