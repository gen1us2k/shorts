package database

import (
	"github.com/gen1us2k/shorts/model"
)

type (
	WriteDatabase interface {
		ShortifyURL(model.URL) (model.URL, error)
		ListURLs(string) ([]model.URL, error)
		GetURLByHash(string) (model.URL, error)
		StoreView(model.Referer) error
	}
	Analytics interface {
		Statistics(model.URL) (model.Statistics, error)
	}
)
