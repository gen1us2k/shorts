package model

import "time"

type (
	URL struct {
		OriginalURL    string
		CreationDate   time.Time
		ExpirationDate time.Time
		UserID         int64
		User           User
	}
	Referer struct {
		URL string
	}
	User struct {
		FirstName    string
		LastName     string
		CreationDate string
		LastLogin    string
		Email        string
	}
	Statistics struct{}
)
