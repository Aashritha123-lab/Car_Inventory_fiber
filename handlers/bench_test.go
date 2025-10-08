package handlers

import (
	"car/config"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func BenchmarkGetCar(b *testing.B) {
	config.ConnectDB()
	app := fiber.New()
	app.Get("/cars/:id", GetCar)
	req, _ := http.NewRequest("GET", "/cars/13", nil)
	req.Header.Set("Content-Type", "application/json")

	// b.ResetTimer()
	for i := 0; i < b.N; i++ {

		_, _ = app.Test(req, 5000)
	}
}
