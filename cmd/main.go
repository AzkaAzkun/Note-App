package main

import (
	"os"

	"github.com/AzkaAzkun/Note-App/config"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	env := os.Getenv("ENV")
	if err != nil && env == " " {
		panic("Error open .Env file")
	}
	config.ConnectDatabase()

	app := fiber.New()
	router.SetupAuth(app)
}
