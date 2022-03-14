package database

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gen1us2k/shorts/config"
	"github.com/gen1us2k/shorts/model"
	"github.com/supabase/postgrest-go"
)

type Supabase struct {
	WriteDatabase
	config *config.ShortsConfig
	conn   *postgrest.Client
}

func NewSupabase(c *config.ShortsConfig) (*Supabase, error) {
	conn := postgrest.NewClient(c.SupabaseURL, "", map[string]string{
		"apikey":        c.SupabaseKey,
		"Authorization": fmt.Sprintf("Bearer %s", c.SupabaseKey),
	})
	if conn.ClientError != nil {
		return nil, conn.ClientError
	}
	return &Supabase{
		config: c,
		conn:   conn,
	}, nil
}

func (s *Supabase) ShortifyURL(url model.URL) (model.URL, error) {
	var urls []model.URL
	url.Hash = generateHash(s.config.URLLength)
	url.ExpiredAt = time.Now().AddDate(0, 1, -1) // FIXME: Make this configurable
	q := s.conn.From("url").Insert(url, false, "do-nothing", "", "")
	_, err := q.ExecuteTo(&urls)
	return urls[0], err
}
func (s *Supabase) ListURLs(ownerID string) ([]model.URL, error) {
	var urls []model.URL
	q := s.conn.From("url").Select("*", "10", false).Match(map[string]string{"owner_id": ownerID}) // FIXME: Make this configurable
	_, err := q.ExecuteTo(&urls)
	return urls, err
}
func (s *Supabase) GetURLByHash(hash string) (model.URL, error) {
	var urls []model.URL
	q := s.conn.From("url").Select("*", "10", false).Match(map[string]string{"hash": hash})
	_, err := q.ExecuteTo(&urls)
	if len(urls) == 0 {
		return model.URL{}, errors.New("does not exist")
	}
	return urls[0], err
}
func (s *Supabase) StoreView(ref model.Referer) error {
	q := s.conn.From("url_view").Insert(ref, false, "do-nothing", "", "")
	_, _, err := q.Execute()
	return err
}
func (s *Supabase) DeleteURL(url model.URL) error {
	q := s.conn.From("url").Delete("", "").Match(map[string]string{"id": strconv.FormatInt(url.ID, 10)})
	_, _, err := q.Execute()
	return err

}
