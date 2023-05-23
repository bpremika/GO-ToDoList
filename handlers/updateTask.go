package handlers

import (
	"fmt"
	"time"

	"github.com/bpremika/go-toDoList/database"
	"github.com/bpremika/go-toDoList/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateTaskSchema struct {
	Description string    `json:"description"`
	Duedate     time.Time `json:"duedate"`
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := models.Task{}
	body := UpdateTaskSchema{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if result := database.DB.Db.First(&task, id); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
	fmt.Printf("body: %v\n", body)
	task.Description = &body.Description
	task.Duedate = &body.Duedate
	database.DB.Db.Save(&task)
	return c.Status(200).JSON(&task)
}
