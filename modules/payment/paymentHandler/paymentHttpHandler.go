package paymentHandler

import (
	"github.com/natdanai0917/test_repo/config"
	"github.com/natdanai0917/test_repo/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{cfg, paymentUsecase}
}
