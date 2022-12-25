package main

import (
	"context"
	"currency_exchange/internal/clients/currate"
	"currency_exchange/internal/config"
	"currency_exchange/internal/service"
	"currency_exchange/internal/storage"
	"currency_exchange/internal/transport/http"
	"currency_exchange/internal/transport/http/handlers"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	//asd := currate.ClientResponse{}
	//asd.Status = 200
	//asd.Message = "asdf"
	//asd.Data = make(map[string]string)
	//asd.Data["USDRUB"] = "1234"
	//
	//data, _ := json.Marshal(asd)
	//fmt.Println(string(data))
	//return

	out := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.StampMilli,
	}

	logger := zerolog.New(out).With().Caller().Logger().With().Timestamp().Logger()

	cfg, err := config.Parse()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse configs")
	}

	dbConnStr := createConnectionString(
		cfg.DataBase.Host,
		cfg.DataBase.Database,
		cfg.DataBase.User,
		cfg.DataBase.Password,
	)
	fmt.Println(dbConnStr)
	pool, err := newPgxPool(dbConnStr)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create pgx pool")
	}

	psqlStorage := storage.New(pool)
	client := currate.New(logger)

	svc := service.New(logger, psqlStorage, client)

	handler := handlers.New(logger, svc)
	fmt.Println(handler.CreatePairHandler)
	server := http.New(cfg.Server.Host, handler)

	go func() {
		if err = server.Run(); err != nil {
			logger.Fatal().Err(err).Msg("server starting error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	//
	//logger.Info().Msg("http server shutdown")
	//
	//if err = server.Shutdown(); err != nil {
	//	logger.Fatal().Err(err).Msg("server shutdown error")
	//}
}

func createConnectionString(host, db, user, password string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		user,
		password,
		host,
		db,
	)
}

func newPgxPool(connString string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать конфиг *pgxpool.Config %v", err)
	}
	pool, err := pgxpool.ConnectConfig(
		context.Background(),
		cfg,
	)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать *pgxpool.Pool %v", err)
	}

	return pool, nil
}
