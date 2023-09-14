package routes

import (
	"github.com/dilson-m/golang-fiber-crud/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/register", controller.Register)
}
