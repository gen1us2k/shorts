package database

import (
	"time"

	"github.com/gen1us2k/shorts/model"
)

type (
	WriteDatabase interface {
		ShortifyURL(model.URL) (model.URL, error)
		GetURLByHash(string) (URL, error)
		StoreView() error
		SaveStatistics() error
	}
	Analytics interface {
		Statistics(model.URL) (model.Statistics, error)
	}
	// URL struct is required for working with databases
	URL struct {
		ID        int64     `db:"id"`
		URL       string    `db:"url"`
		Hash      string    `db:"hash"`
		CreatedAt time.Time `db:"created_at"`
		ExpiredAt time.Time `db:"expired_at"`
		OwnerID   string    `db:"owner_id"`
	}
)
