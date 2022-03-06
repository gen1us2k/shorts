package api

import (
	"log"
	"net/http"

	"github.com/gen1us2k/shorts/config"
	"github.com/gen1us2k/shorts/database"
	"github.com/gen1us2k/shorts/model"
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
	s := &Server{
		r:      gin.Default(),
		config: c,
		db:     db,
	}
	s.initRoutes()
	return s, nil
}
func (s *Server) initRoutes() {
	s.r.GET("/:hash", s.showURL)
	// The only kratos thing would be here
	userAPI := s.r.Group("/api/")
	userAPI.POST("/url", s.shortifyURL)
	userAPI.GET("/url", s.listURLs)

	// TODO: Implement RBAC here
	analyticsAPI := s.r.Group("/analytics")
	_ = analyticsAPI

}

func (s *Server) showURL(c *gin.Context) {
	log.Println("showURL")
	hash := c.Param("hash")
	u, err := s.db.GetURLByHash(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "error querying database",
		})
		return
	}
	ref := model.Referer{
		URLID:   u.ID,
		Referer: c.Request.Header["Referer"][0],
	}
	if err := s.db.StoreView(ref); err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "error querying database",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, u.URL)
}
func (s *Server) listURLs(c *gin.Context) {
	log.Println("listURLS")
	// FIXME: Improve handling
	urls, err := s.db.ListURLs("something")
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "error querying database",
		})
		return
	}
	c.JSON(http.StatusOK, urls)
}

func (s *Server) shortifyURL(c *gin.Context) {
	log.Println("Shortify")
	var url model.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, &model.DefaultResponse{
			Message: "failed parse json",
		})
		return
	}
	u, err := s.db.ShortifyURL(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.DefaultResponse{
			Message: "failed to create short version",
		})
		return
	}
	c.JSON(http.StatusOK, u)
}

func (s *Server) Start() error {
	return s.r.Run(s.config.BindAddr)
}
