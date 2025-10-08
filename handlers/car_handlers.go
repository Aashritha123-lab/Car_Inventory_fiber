package handlers

import (
	"car/models"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// to make concurrent operation seamlessly we use mutexes
var Mu sync.Mutex

func CreateCar(c *fiber.Ctx) error {
	Mu.Lock()
	defer Mu.Unlock()

	car := &models.Car{}
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Incorrect input body",
			"detail": err.Error(),
		})
	}
	car.Insert()

	fmt.Println("Car saved to the inventory with id", car.ID)
	return c.Status(fiber.StatusCreated).JSON(car)

}

func GetCar(c *fiber.Ctx) error {
	Mu.Lock()
	defer Mu.Unlock()

	car := &models.Car{}

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "invalid car id",
			"detail": err.Error(),
		})
	}
	car.ID = id
	if err = car.Get(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "car not found with id",
			"id":    car.ID,
		})
	}

	// fmt.Println("Found the car:", http.StatusFound)
	return c.Status(fiber.StatusFound).JSON(car)
}

func DeleteCar(c *fiber.Ctx) error {
	Mu.Lock()
	defer Mu.Unlock()

	car := &models.Car{}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid car id",
			"id":    id,
		})
	}
	car.ID = id
	if err := car.Delete(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Car not found with id so nothing deleted",
			"id":    car.ID,
		})
	}

	fmt.Println("Deleted the car :", http.StatusOK)

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateCar(c *fiber.Ctx) error {
	Mu.Lock()
	defer Mu.Unlock()

	car := &models.Car{}

	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Invalid input body",
			"detail": err.Error(),
		})
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid car id",
			"id":    id,
		})
	}

	car.ID = id

	if err := car.Update(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Car not found with id so nothing updated",
			"id":    car.ID,
		})
	}

	fmt.Println("Car updated to the inventory with id", car.ID)

	return c.Status(fiber.StatusOK).JSON(car)
}
