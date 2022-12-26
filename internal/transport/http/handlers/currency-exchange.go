package handlers

import (
	"currency_exchange/internal/models"
	"github.com/gofiber/fiber"
	"net/http"
)

func (h *Handler) Exchange(c *fiber.Ctx) {
	var pair *models.CurrencyPair
	if err := c.BodyParser(&pair); err != nil {
		c.JSON(models.Response{
			Error:     true,
			ErrorText: err.Error(),
			Code:      http.StatusBadRequest,
		})
		return
	}

	total, err := h.service.CurrencyExchange(c.Context(), pair)
	if err != nil {
		c.JSON(models.Response{
			Error:     true,
			ErrorText: err.Error(),
			Code:      http.StatusInternalServerError,
		})
		return
	}

	c.JSON(models.Response{
		Data: total,
		Code: http.StatusOK,
	})
}
