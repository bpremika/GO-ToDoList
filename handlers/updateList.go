package handlers

import (
	"github.com/bpremika/go-toDoList/database"
	"github.com/bpremika/go-toDoList/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateListSchema struct {
	Title string `json:"title"`
}

func UpdateList(c *fiber.Ctx) error {
	id := c.Params("id")
	list := models.List{}
	body := UpdateListSchema{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if result := database.DB.Db.First(&list, id); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	list.Title = body.Title
	database.DB.Db.Save(&list)
	return c.Status(200).JSON(&list)
}
