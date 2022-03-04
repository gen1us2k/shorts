package database

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gen1us2k/shorts/config"
	"github.com/gen1us2k/shorts/model"
)

type Postgres struct {
	WriteDatabase
	conn *sql.DB
}

func NewPostgres(c *config.ShortsConfig) (*Postgres, error) {
	db, err := sql.Open("postgres", c.DSN)
	if err != nil {
		return nil, err
	}
	return &Postgres{conn: db}, nil
}
func (p *Postgres) ShortifyURL(u model.URL) (model.URL, error) {
	return model.URL{}, nil
}
func (p *Postgres) StoreView() error { return nil }
