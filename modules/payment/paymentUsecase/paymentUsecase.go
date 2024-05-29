package paymentUsecase

import (
	"github.com/natdanai0917/test_repo/modules/payment/paymentRepository"
)

type (
	PaymentUsecaseService interface{}

	paymentUsecase struct {
		paymentRepository paymentRepository.PaymentRepositoryService
	}
)

func NewPaymentUseCase(paymentRepository paymentRepository.PaymentRepositoryService) PaymentUsecaseService {
	return &paymentUsecase{paymentRepository}
}
