package main

import (
	"car/config"
	"car/handlers"
	"car/middleware"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.ConnectDB()

	app := fiber.New()

	app.Use(logger.New())

	app.Use(middleware.SecurityHeader)
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "admin123",
			"ashri": "ashri123",
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized access to the car inventory",
			})
		},
	}))

	// POST : /cars
	// GET : /cars/8091
	// DELETE : /cars/8091
	// PUT : /cars/8091

	app.Post("/cars", handlers.CreateCar)
	app.Get("/cars/:id", handlers.GetCar)
	app.Delete("/cars/:id", handlers.DeleteCar)
	app.Put("/cars/:id", handlers.UpdateCar)

	fmt.Println("Fiber HTTP server is  listening on port 3051....")
	if err := app.Listen(":3051"); err != nil {
		log.Fatalf("Error in starting the server : %v", err)
	}
}
