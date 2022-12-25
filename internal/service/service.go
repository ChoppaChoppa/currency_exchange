package service

import (
	"context"
	"currency_exchange/internal/clients/currate"
	"currency_exchange/internal/models"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type IStorage interface {
	GetAllPair(ctx context.Context) ([]*models.CurrencyPair, error)
	GetExchangeRate(ctx context.Context, FromCurrency, ToCurrency string) (float64, error)
	CreateCurrencyPair(ctx context.Context, pair *models.CurrencyPair) error
}

type ICurrateClient interface {
	Get(from, to string) (*currate.ClientResponse, error)
}

type service struct {
	logger  zerolog.Logger
	storage IStorage
	client  ICurrateClient
}

func New(logger zerolog.Logger, storage IStorage, client ICurrateClient) *service {
	return &service{
		logger:  logger,
		storage: storage,
		client:  client,
	}
}

func (s *service) CreatePair(ctx context.Context, pair *models.CurrencyPair) error {
	resp, err := s.client.Get(pair.From, pair.To)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to get resp from client")
		return err
	}

	if resp.Status != http.StatusOK {
		return ErrBadRequest
	}

	for _, v := range resp.Data {
		well, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return ErrServer
		}

		pair.Well = well
	}

	if err = s.storage.CreateCurrencyPair(ctx, pair); err != nil {
		s.logger.Error().Err(err).Msg("failed to create curr pair db")
		return err
	}

	return nil
}
