package api

import (
	"github.com/gen1us2k/shorts/config"
	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		r      gin.Router
		config *config.ShortsConfig
	}
)

func New(c *config.ShortsConfig) (*Server, error) {
	return &Server{
		r:      gin.Default(),
		config: c,
	}, nil
}

func (s *Server) initRoutes() {
	// The only kratos thing would be here
	userAPI := s.r.Group("/api/")
	_ = userAPI

	// TODO: Implement RBAC here
	analyticsAPI := s.r.Group("/analytics")
	_ = analyticsAPI

}

func (s *Server) Start() error {
	return s.r.Run(s.config.BindAddr)
}
