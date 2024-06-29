package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString(os.Getenv("APP_NAME"))
	})

	// load .env file if not in production
	if env := os.Getenv("RAILWAY_ENVIRONMENT_NAME"); env != "production" {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
