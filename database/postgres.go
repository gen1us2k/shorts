package database

import (
	_ "github.com/lib/pq"

	"github.com/gen1us2k/shorts/config"
	"github.com/gen1us2k/shorts/model"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	WriteDatabase
	config *config.ShortsConfig
	conn   *sqlx.DB
}

func NewPostgres(c *config.ShortsConfig) (*Postgres, error) {
	db, err := sqlx.Open("postgres", c.DSN)
	if err != nil {
		return nil, err
	}
	return &Postgres{conn: db, config: c}, nil
}
func (p *Postgres) ShortifyURL(u model.URL) (model.URL, error) {
	var url model.URL
	tx, err := p.conn.Begin()
	if err != nil {
		return url, err
	}
	u.Hash = generateHash(p.config.URLLength)
	err = tx.Get(&url, "INSERT INTO url (url, hash, expired_at, owner_id) VALUES($1, $2, $3, $4) RETURNING (id, url, hash, created_at, expired_at, owner_id)")
	if err != nil {
		tx.Rollback()
		return url, err
	}
	tx.Commit()

	return url, nil
}
func (p *Postgres) ListURLs(ownerID string) ([]model.URL, error) {
	var urls []model.URL
	err := p.conn.Select(&urls, "SELECT * FROM url WHERE owner_id=$1", ownerID)
	return urls, err
}

func (p *Postgres) StoreView(ref model.Referer) error {
	tx, err := p.conn.Begin()
	if err != nil {
		return err
	}
	err = tx.Exec("INSERT INTO url_view (referer, url_id) VALUES ($1, $2)", ref.Referer, ref.URLID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
