package database

import "github.com/gen1us2k/shorts/model"

type (
	WriteDatabase interface {
		ShortifyURL(model.URL) (model.URL, error)
		StoreView() error
	}
	Analytics interface {
		Statistics(model.URL) (model.Statistics, error)
	}
)
