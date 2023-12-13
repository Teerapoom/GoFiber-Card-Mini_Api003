package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teerapoom/mini_api003/schema"
)

func GetUser(c *fiber.Ctx) error {
	var users []schema.CreditCard
	err := schema.GetUser(&users)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.JSON(users)
}
