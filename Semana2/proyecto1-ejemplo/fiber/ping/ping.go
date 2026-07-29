package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const TargetIP = "192.168.122.133"

func main() {
	app := fiber.New()

	// Middleware de logger para ver las peticiones en consola
	app.Use(logger.New())

	// Endpoint /ping migrado a Fiber
	app.Get("/ping", func(c *fiber.Ctx) error {
		resp, err := http.Get("http://" + TargetIP + ":8082/pong")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error contactando a Pong: %s", err))
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error leyendo respuesta de Pong")
		}

		return c.SendString(fmt.Sprintf("Ping: Llame a la otra API y me dijo: %s", string(body)))
	})

	// Endpoint de salud (Health Check)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"status":  "UP",
			"message": "Ping API (Fiber) is Ready",
		})
	})

	fmt.Println("API PING corriendo en puerto 8081")
	log.Fatal(app.Listen(":8081"))
}
