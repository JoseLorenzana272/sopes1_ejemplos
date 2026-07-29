package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Middleware de logger para ver las peticiones
	app.Use(logger.New())

	// Endpoint /pong migrado a Fiber
	app.Get("/pong", func(c *fiber.Ctx) error {
		return c.SendString("¡Pong! (Desde Containerd en VM 2)")
	})

	// Endpoint de salud (Health Check)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"status":  "UP",
			"message": "Pong API (Fiber) is Ready",
		})
	})

	fmt.Println("API PONG corriendo en puerto 8082")
	log.Fatal(app.Listen(":8082"))
}
