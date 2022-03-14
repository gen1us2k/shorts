package database

import (
	"github.com/gen1us2k/shorts/config"
	"github.com/gen1us2k/shorts/model"
)

type (
	// WriteDatabase interface represents interface for
	// a database layer with write permissions. For Analytics you need to implement
	// Analytics interface
	WriteDatabase interface {
		ShortifyURL(model.URL) (model.URL, error)
		ListURLs(string) ([]model.URL, error)
		GetURLByHash(string) (model.URL, error)
		StoreView(model.Referer) error
		DeleteURL(model.URL) error
	}
	// Analytics interface required for analytics
	// for Write API you need to implement WriteDatabase interface
	Analytics interface {
		Statistics(model.URL) (model.Statistics, error)
	}
)

func CreateStorage(c *config.ShortsConfig) (WriteDatabase, error) {
	if c.DatabaseProvider == config.ProviderPostgres {
		return NewPostgres(c)
	}
	if c.DatabaseProvider == config.ProviderSupabase {
		return NewSupabase(c)
	}
	return NewPostgres(c)
}
