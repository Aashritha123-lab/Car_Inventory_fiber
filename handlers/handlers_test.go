package handlers

import (
	"car/config"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

//github.com/stretchr/testify/assert

// tests := []struct {
// 	name           string
// 	body           string
// 	expectedStatus int
// 	expectedbody   string
// }{
// 	{
// 		name:           "Create Car - Valid Input",
// 		body:           `{"Name":"X7","Model":"V8","Brand":"BMW","Year":2023,"Price":20000000}`,
// 		expectedStatus: http.StatusCreated,
// 		expectedbody:   `"Name":"X7"`,
// 	},
// 	{
// 		name:           "Invalid body",
// 		body:           `{"Name":"X7","Model":"V8","Brand":"BMW","Year":2023}`,
// 		expectedStatus: http.StatusBadRequest,
// 		expectedbody:   `Input Json format is wrong`,
// 	},
// 	{
// 		name:           "Invalid request",
// 		body:           ``,
// 		expectedStatus: http.StatusBadRequest,
// 		expectedbody:   `Input Json format is wrong`,
// 	},
// }

// func TestCreateCarInvalidInput(t *testing.T) {
// 	config.ConnectDB()

// 	app := fiber.New()

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			app.Post("/cars", CreateCar)
// 			req, _ := http.NewRequest("POST", "/cars", strings.NewReader(tt.body))
// 			req.Header.Set("Content-Type", "application/json")
// 			resp, err := app.Test(req, 5000)
// 			if err != nil {
// 				t.Fatalf("Error making request: %v", err)
// 			}
// 			bodyBytes,_ := io.ReadAll(resp.Body)
// 			bodyString := string(bodyBytes)
// 			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
// 			assert.Contains(t, bodyString, tt.expectedbody)
// 		})
// 	}

// }

func TestCreateCar(t *testing.T) {
	config.ConnectDB()

	app := fiber.New()
	app.Post("/cars", CreateCar)

	body := `{"Name":"X7","Model":"V8","Brand":"BMW","Year":2023,"Price":20000000}`

	req, _ := http.NewRequest("POST", "/cars", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

//

func TestGetCar(t *testing.T) {
	// config.ConnectDB()

	app := fiber.New()
	app.Get("/cars/:id", GetCar)

	req, _ := http.NewRequest("GET", "/cars/13", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	assert.Equal(t, http.StatusFound, resp.StatusCode)
}
