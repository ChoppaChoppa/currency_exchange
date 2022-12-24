package main

import (
	"currency_exchange/internal/config"
	"currency_exchange/internal/transport/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	out := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.StampMilli,
	}

	logger := zerolog.New(out).With().Caller().Logger().With().Timestamp().Logger()

	cfg, err := config.Parse()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse configs")
	}

	server := http.New(cfg.Server.Host)

	go func() {
		if err = server.Run(); err != nil {
			logger.Fatal().Err(err).Msg("server starting error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info().Msg("http server shutdown")

	if err = server.Shutdown(); err != nil {
		logger.Fatal().Err(err).Msg("server shutdown error")
	}
}
