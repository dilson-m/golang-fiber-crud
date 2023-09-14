package controller

import (
	"github.com/dilson-m/golang-fiber-crud/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	user := models.User{
		FirstName: "Dilson",
	}

	user.LastName = "Mascarenhas"

	return c.JSON(user)
	// return c.SendString("Bem vindo")
}
