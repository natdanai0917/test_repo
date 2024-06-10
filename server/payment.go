package server

import (
	"github.com/natdanai0917/test_repo/modules/payment/paymentHandler"
	"github.com/natdanai0917/test_repo/modules/payment/paymentRepository"
	"github.com/natdanai0917/test_repo/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	paymentRepository := paymentRepository.NewPaymentRepository(s.db)
	paymentUsecase := paymentUsecase.NewPaymentUseCase(paymentRepository)
	paymentHttpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, paymentUsecase)

	paymentQueueHandler := paymentHandler.NewPaymentQueueHandler(s.cfg, paymentHttpHandler)

	_ = paymentHttpHandler
	_ = paymentQueueHandler

	payment := s.app.Group("/payment_v1")

	//Health Check
	payment.GET("", s.healthCheckService)
}
