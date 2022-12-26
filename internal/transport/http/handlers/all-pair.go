package handlers

import (
	"currency_exchange/internal/models"
	"github.com/gofiber/fiber"
)

func (h *Handler) GetAllPairs(c *fiber.Ctx) {
	pairs, err := h.service.GetAllPairs(c.Context())
	if err != nil {
		c.JSON(models.Response{
			Error:     true,
			ErrorText: err.Error(),
		})
		return
	}

	c.JSON(models.Response{
		Data: pairs,
	})
}
