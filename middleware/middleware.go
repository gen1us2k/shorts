package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gen1us2k/shorts/config"
	"github.com/gin-gonic/gin"
)

type (
	KratosMiddleware struct {
		config *config.ShortsConfig
		client *http.Client
	}
	// Identity represents identity sent from Kratos
	Identity struct {
		ID     string `json:"id"`
		Traits struct {
			Name struct {
				First string `json:"first"`
				Last  string `json:"last"`
			} `json:"name"`
			Email string `json:"email"`
		} `json:"traits"`
	}
	// KratosSession represents Kratos session returned
	// from /session/whoami from Ory Kratos
	KratosSession struct {
		Active                bool      `json:"active"`
		AuthenticatedAt       time.Time `json:"authenticated_at"`
		AuthenticationMethods []struct {
			CompletedAt time.Time `json:"completed_at"`
			Method      string    `json:"method"`
		} `json:"authentication_methods"`
		AuthenticatorAssuranceLevel string   `json:"authenticator_assurance_level"`
		ExpiresAt                   string   `json:"expires_at"`
		Identity                    Identity `json:"identity"`
		ID                          string   `json:"id"`
	}
)

func NewAuthenticationMiddleware(c *config.ShortsConfig) *KratosMiddleware {
	return &KratosMiddleware{
		config: c,
		client: &http.Client{},
	}
}
func (k *KratosMiddleware) AuthenticationMiddleware(c *gin.Context) {
	session, err := k.validateSession(c.Request)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, k.config.UIURL)
		return
	}
	c.Set(config.OwnerKey, session.Identity.ID)
	c.Next()
}
func (k *KratosMiddleware) validateSession(r *http.Request) (*KratosSession, error) {
	var session KratosSession
	cookie, err := r.Cookie(config.KratosSessionKey)
	if err != nil {
		return nil, err
	}
	if cookie == nil {
		return nil, errors.New("no session in cookies")
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/sessions/whoami", k.config.APIURL), http.NoBody)
	if err != nil {
		return nil, err
	}
	req.AddCookie(cookie)
	resp, err := k.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("wrong status code")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, err
	}
	return &session, nil
}
