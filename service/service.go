package service

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"testCDN/handler"
	"time"
)

type Service struct {
	router  *gin.Engine
	handler *handler.Handler
	logger  *zap.SugaredLogger
}

func NewService(logger *zap.SugaredLogger) *Service {
	router := gin.New()
	if gin.IsDebugging() {
		router.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
	}
	router.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))
	h := handler.NewHandler(logger)
	return &Service{
		router:  router,
		handler: h,
		logger:  logger,
	}
}

func (s *Service) Start() {
	s.SetRouter()
	if err := s.router.Run(":8080"); err != nil {
		s.logger.Fatalf("Service start: router run: %v", err)
	}
}
