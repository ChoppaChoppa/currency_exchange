package http

import (
	"currency_exchange/internal/transport/http/handlers"
	"github.com/gofiber/fiber"
)

type Server struct {
	*fiber.App
	host     string
	handlers *handlers.Handler
}

func New(host string, handlers *handlers.Handler) *Server {
	return &Server{
		App:      fiber.New(),
		host:     host,
		handlers: handlers,
	}
}

func (s *Server) Run() error {
	InitRouters(s)

	if err := s.App.Listen(":3030"); err != nil {
		return err
	}

	return nil
}
