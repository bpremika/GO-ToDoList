package handlers

import (
	"github.com/bpremika/go-toDoList/database"
	"github.com/bpremika/go-toDoList/models"
	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	tasks := []models.Task{}
	results := database.DB.Db.Order("position").Find(&tasks)
	if results.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": results.Error})
	}
	return c.Status(200).JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	newTask := &models.Task{}
	if err := c.BodyParser(newTask); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&newTask)

	return c.Status(200).JSON(newTask)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := models.Task{}
	result := database.DB.Db.Delete(&task, id)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
	return c.Status(200).JSON(task)
}
