package middlewareUsecase

import "github.com/natdanai0917/test_repo/modules/middleware/middlewareRepository"

type (
	MiddlewareUsecaseService interface{}

	middlewareUsecase struct {
		MiddlewareRepositoryService middlewareRepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepository}
}
