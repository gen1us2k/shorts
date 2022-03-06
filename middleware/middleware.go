package middleware

import (
	"github.com/gen1us2k/shorts/config"
	"github.com/gin-gonic/gin"
)

type KratosMiddleware struct {
	config *config.ShortsConfig
}

func (k *KratosMiddleware) User(c *gin.Context) {
}
