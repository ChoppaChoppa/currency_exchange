package handlers

import (
	"context"
	"currency_exchange/internal/models"
	"github.com/rs/zerolog"
)

type IService interface {
	CreatePair(ctx context.Context, pair *models.CurrencyPair) error
}

type Handler struct {
	logger  zerolog.Logger
	service IService
}

func New(logger zerolog.Logger, service IService) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}
