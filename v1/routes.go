package v1

import (
	"github.com/gofiber/fiber/v2"
)

// Routes - Configuração das rotas
func Routes() *fiber.App {
	app := fiber.New(fiber.Config{})
	app.Get("/exchange/:amount/:from/:to", GetCounting)
	return app
}
