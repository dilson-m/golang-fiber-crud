package controller

import (
	"github.com/dilson-m/golang-fiber-crud/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	if data["Password"] != data["Confirm_Password"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "As senhas nao sao iguais!",
		})
	}

	user := models.User{
		FirstName: data["First_name"],
		LastName:  data["Last_name"],
		Email:     data["Email"],
		Password:  data["Password"],
	}

	return c.JSON(user)
	// return c.SendString("Bem vindo")
}
