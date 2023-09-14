package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// fmt.Println("Bem vindo ao crud em golang com FIBER.")

	//	created instance FIBER
	app := fiber.New()

	// created httpHandler
	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Bem Vindo!")
		return c.Status(200).JSON(fiber.Map{
			"success":true,
			"message":"Go FIBER API created with success!",
		})
	})

	// listen on portapp.Listen(":3000")

}
