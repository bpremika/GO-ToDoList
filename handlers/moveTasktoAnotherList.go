package handlers

import (
	"github.com/bpremika/go-toDoList/database"
	"github.com/bpremika/go-toDoList/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateTaskList struct {
    ListID int `json:"list_id"`
}

func MoveTasktoAnotherList(c *fiber.Ctx) error {
	id := c.Params("id")
	task := models.Task{}	
	body := UpdateTaskList{}

	if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }
	if result := database.DB.Db.First(&task, id); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
	task.ListID = body.ListID
	database.DB.Db.Save(&task)

	return c.Status(200).JSON(&task)
}
