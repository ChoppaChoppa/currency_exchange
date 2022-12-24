package models

import "time"

var (
	USD Currency = "USD"
	RUB Currency = "RUB"
	BTC Currency = "BTC"
	EUR Currency = "EUR"
)

type Currency string

type CurrencyPair struct {
	ID        int       `json:"id"`
	Well      int       `json:"well"`
	From      Currency  `json:"from"`
	To        Currency  `json:"to"`
	UpdatedAt time.Time `json:"updated_at"`
}
