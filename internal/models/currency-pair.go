package models

import "time"

type CurrencyPair struct {
	ID        int       `json:"id" db:"id"`
	Well      float64   `json:"well" db:"well"`
	From      string    `json:"currency_from" db:"currency_from"`
	To        string    `json:"currency_to" db:"currency_to"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
