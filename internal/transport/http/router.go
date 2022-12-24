package http

func InitRouters(s *Server) {
	api := s.Group("/api")

	api.Post("/currency", nil)
	api.Get("/currency", nil)
	api.Get("currency", nil)
}
