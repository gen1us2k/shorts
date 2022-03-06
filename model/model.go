package model

import "time"

type (
	// URL struct is required for working with databases
	URL struct {
		ID        int64     `db:"id"`
		URL       string    `db:"url"`
		Hash      string    `db:"hash"`
		CreatedAt time.Time `db:"created_at"`
		ExpiredAt time.Time `db:"expired_at"`
		OwnerID   string    `db:"owner_id"`
	}
	Referer struct {
		ID      int64  `db:"id"`
		URLID   int64  `db:"url_id"`
		Referer string `db:"referer"`
	}
	User struct {
		FirstName    string
		LastName     string
		CreationDate string
		LastLogin    string
		Email        string
	}
	DefaultResponse struct {
		Message string `json:"message"`
	}
	Statistics struct{}
)
