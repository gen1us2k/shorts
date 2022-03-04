package api

import (
	"github.com/gen1us2k/shorts/config"
	"github.com/gen1us2k/shorts/database"
	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		r      *gin.Engine
		config *config.ShortsConfig
		db     database.WriteDatabase
	}
)

func New(c *config.ShortsConfig) (*Server, error) {
	db, err := database.NewPostgres(c)
	if err != nil {
		return nil, err
	}
	return &Server{
		r:      gin.Default(),
		config: c,
		db:     db,
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
