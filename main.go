package main

import (
	"github.com/dilson-m/golang-fiber-crud/database"
	"github.com/dilson-m/golang-fiber-crud/routes"
	"github.com/gofiber/fiber/v2"
)

// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("nao pode dividir por zero")
// 	}

// 	return a / b, nil

// }

func main() {
	// d, e := divide(2, 2)
	// fmt.Printf("d:%v\ne:%v\n", d, e)

	//	faz conexao com banco de dados
	database.Connect()
	// Cria uma instancia
	app := fiber.New()
	routes.Setup(app)

	// Define pagina inicial + Handle
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{
	// 		"success": true,
	// 		"message": "Conexao Ok.",
	// 	})
	// })

	app.Listen(":3000")

}
