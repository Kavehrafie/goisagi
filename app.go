package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString(os.Getenv("APP_NAME"))
	})

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	log.Fatal(app.Listen(":3000"))
}
