package handlers

import (
	"currency_exchange/internal/models"
	"fmt"
	"github.com/gofiber/fiber"
	"net/http"
)

func (h *Handler) CreatePairHandler(c *fiber.Ctx) {
	var pair *models.CurrencyPair
	if err := c.BodyParser(&pair); err != nil {
		h.logger.Error().Err(err).Msg("failed to parse body")
		c.JSON(models.Response{
			Error:     true,
			ErrorText: err.Error(),
			Code:      http.StatusInternalServerError,
		})
		return
	}
	fmt.Println(pair.From, pair.To)
	if err := h.service.CreatePair(c.Context(), pair); err != nil {
		h.logger.Error().Err(err).Msg("failed to create pair")
		c.JSON(models.Response{
			Error:     true,
			ErrorText: err.Error(),
			Code:      http.StatusInternalServerError,
		})
		return
	}

	c.JSON(models.Response{
		Code: http.StatusOK,
	})
}
