package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teerapoom/mini_api003/schema"
)

func CreateCreditCard(c *fiber.Ctx) error {
	var input schema.CreditCard
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	card := schema.CreditCard{
		Fullname: input.Fullname,
		Number:   input.Number,
	}

	_, err := card.SaveCard()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(card)
}
