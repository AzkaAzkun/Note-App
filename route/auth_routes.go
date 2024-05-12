package route

import (
	"github.com/AzkaAzkun/Note-App/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupAuth(app *fiber.App) {
	app.Post("/login", handler.Login)
}
