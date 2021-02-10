package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) SetRouter() {
	s.router.GET("/", s.health)
	s.router.GET("/health", s.health)
	s.router.GET("/sortFile", s.handler.SortFile)
	s.router.POST("/sort", s.handler.Sort)
}

func (s *Service) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}
