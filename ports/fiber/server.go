package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const defaultPort = "3000"

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	srv := fiber.New()

	srv.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Stockroom port(%v) ... welcome!", os.Getenv("FIBER_PORT")))
	})

	port := os.Getenv("FIBER_PORT")
	if port == "" {
		port = defaultPort
	}
	log.Fatal(srv.Listen(fmt.Sprintf(":%s", port)))
}
