package route

import (
	"testgoapi/controller/wallet"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) *fiber.App {

	// index
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// wallet
	app.Put("/api/update_balance", wallet.UpdateBalance)

	return app
}
