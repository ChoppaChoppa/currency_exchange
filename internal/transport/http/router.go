package http

import (
	"fmt"
	"github.com/gofiber/fiber/middleware"
)

func InitRouters(s *Server) {
	api := s.App.Group("/api")
	api.Use(middleware.Logger())

	fmt.Println("PI", api)
	api.Post("/currency", s.handlers.CreatePairHandler)
	api.Get("/currency", nil)
	api.Get("currency", nil)
}
