package storage

import (
	"context"
	"currency_exchange/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type storage struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *storage {
	return &storage{
		pool: pool,
	}
}

func (s *storage) CreateCurrencyPair(ctx context.Context, pair *models.CurrencyPair) error {
	execQuery := `
		INSERT INTO currency (currency_from, currency_to, well)
		VALUES ($1, $2, $3);
	`
	getQuery := `
		SELECT id
		FROM currency
		WHERE currency_from = $1 AND currency_to = $2
	`

	var id int
	err := s.pool.QueryRow(
		ctx,
		getQuery,
		pair.From,
		pair.To,
	).Scan(&id)
	if err == nil {
		return ErrPairExist
	}

	_, err = s.pool.Exec(
		ctx,
		execQuery,
		pair.From,
		pair.To,
		pair.Well,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) GetExchangeRate(ctx context.Context, FromCurrency, ToCurrency string) (float64, error) {
	query := `
		SELECT well
		FROM currency
		WHERE currency_from = $1 AND currency_to = $2
	`

	var well float64
	err := s.pool.QueryRow(
		ctx,
		query,
		FromCurrency,
		ToCurrency,
	).Scan(&well)
	if err != nil {
		return 0, nil
	}

	return well, nil
}

func (s *storage) GetAllPair(ctx context.Context) ([]*models.CurrencyPair, error) {
	query := `
		SELECT id, currency_from, currency_to
		FROM currency
	`

	pairs := make([]*models.CurrencyPair, 0)

	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pair := new(models.CurrencyPair)

		if err = rows.Scan(
			&pair.ID,
			&pair.From,
			&pair.To,
		); err != nil {
			return nil, err
		}

		pairs = append(pairs, pair)
	}

	return pairs, nil
}
