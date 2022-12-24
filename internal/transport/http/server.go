package http

import "github.com/gofiber/fiber"

type Server struct {
	*fiber.App
	host string
}

func New(host string) *Server {
	return &Server{
		host: host,
	}
}

func (s *Server) Run() error {
	if err := s.App.Listen(":3030"); err != nil {
		return err
	}

	return nil
}
