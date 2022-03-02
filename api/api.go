package api

import (
	"github.com/gen1us2k/shorts/config"
	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		r      *gin.Engine
		config *config.ShortsConfig
	}
)

func New(c *config.ShortsConfig) (*Server, error) {
	return &Server{
		r:      gin.Default(),
		config: c,
	}, nil
}
func (s *Server) showURL(c *gin.Context) {}
func (s *Server) initRoutes() {
	s.r.GET("/:url", s.showURL)
	// The only kratos thing would be here
	userAPI := s.r.Group("/api/")
	userAPI.POST("/url", s.shortifyURL)
	_ = userAPI

	// TODO: Implement RBAC here
	analyticsAPI := s.r.Group("/analytics")
	_ = analyticsAPI

}
func (s *Server) shortifyURL(c *gin.Context) {}
func (s *Server) Start() error {
	return s.r.Run(s.config.BindAddr)
}
