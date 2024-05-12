package route

import "github.com/gofiber/fiber"

func SetupAuth(app *fiber.App) {
	app.Post("/login", handler.Login)
}
