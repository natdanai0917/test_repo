package server

import (
	"context"
	"github.com/natdanai0917/test_repo/config"
	"github.com/natdanai0917/test_repo/modules/middleware/middlewareHandler"
	"github.com/natdanai0917/test_repo/modules/middleware/middlewareRepository"
	"github.com/natdanai0917/test_repo/modules/middleware/middlewareUsecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		app        *echo.Echo
		db         *mongo.Client
		cfg        *config.Config
		middleware middlewareHandler.MiddlewareUsecaseService
	}
)

func newMiddleware(cfg *config.Config) middlewareHandler.MiddlewareUsecaseService {
	repo := middlewareRepository.NewMiddlewareRepository()
	usecase := middlewareUsecase.NewMiddlewareUsecase(repo)
	return middlewareHandler.NewMiddlewareHandler(cfg, usecase)
}

func (s *server) gracefullShutdown(pctx context.Context, quit <-chan os.Signal) {
	log.Printf("Start service: %s", s.cfg.App.Name)

	<-quit
	log.Printf("Shutting down service: %s", s.cfg.App.Name)

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second) //if App not shut down in 10 secs this statement will shut it down using echo's (Graceful) Shutdown function
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (s *server) httpListening() {
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}
}

func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	s := &server{app: echo.New(), db: db, cfg: cfg, middleware: newMiddleware(cfg)}

	//Basic middleware
	//Request Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error Request time out",
		Timeout:      30 * time.Second,
	}))

	//CORS
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	//Body Limit
	s.app.Use(middleware.BodyLimit("10M"))

	switch s.cfg.App.Name {
	case "auth":
		s.authService()
	case "player":
		s.playerService()
	case "item":
		s.itemService()
	case "inventory":
		s.inventoryService()
	case "payment":
		s.paymentService()
	}

	//Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s.app.Use(middleware.Logger())

	go s.gracefullShutdown(pctx, quit)

	// Listening
	s.httpListening()
}
