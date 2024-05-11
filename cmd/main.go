package main

import (
	"os"

	"github.com/jeypc/go-jwt-mux/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	env := os.Getenv("ENV")
	if err != nil && env == " " {
		panic("Error open .Env file")
	}
	config.ConnectDatabase()

}
