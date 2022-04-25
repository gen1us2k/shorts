package middleware

import (
	"net/http"

	"github.com/gen1us2k/shorts/config"
	"github.com/gin-gonic/gin"
	client "github.com/ory/kratos-client-go"
)

type kratosMiddleware struct {
	client *client.APIClient
	conf   *config.ShortsConfig
}

func NewMiddleware(c *config.ShortsConfig) *kratosMiddleware {
	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: c.KratosAPIURL,
		},
	}
	return &kratosMiddleware{
		client: client.NewAPIClient(configuration),
	}
}
func (k *kratosMiddleware) Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := k.validateSession(c.Request)
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, k.conf.UIURL)
			return
		}
		if !*session.Active {
			// FIXME: Redirect an user to a better place
			c.Redirect(http.StatusMovedPermanently, k.conf.UIURL)
			return
		}
		c.Set(config.OwnerKey, session.Identity.Id)
		c.Next()
	}
}
func (k *kratosMiddleware) validateSession(r *http.Request) (*client.Session, error) {
	cookies := r.Header.Get("Cookie")
	resp, _, err := k.client.V0alpha2Api.ToSession(r.Context()).
		Cookie(cookies).
		Execute()
	return resp, err
}
