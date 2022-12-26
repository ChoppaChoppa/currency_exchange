package currate

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

const (
	apiKey = "5d85126e1e249d64475fd15be2db5c2c"
	host   = "https://currate.ru"
)

// https://currate.ru/api/?get=rates&pairs=USDRUB&key=5d85126e1e249d64475fd15be2db5c2c
type Currate struct {
	*http.Client
	logger zerolog.Logger
}

func New(logger zerolog.Logger) *Currate {
	return &Currate{
		Client: new(http.Client),
		logger: logger,
	}
}

func (c *Currate) Get(from, to string) (*ClientResponse, error) {
	url := parseUrl(from, to)

	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pair *ClientResponse
	if err = json.Unmarshal(buf, &pair); err != nil {
		return nil, err
	}

	return pair, nil
}

func parseUrl(from, to string) string {
	pair := from + to
	return fmt.Sprintf("%s/api/?get=rates&pairs=%s&key=%s", host, pair, apiKey)
}
