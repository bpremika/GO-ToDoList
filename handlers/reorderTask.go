package handlers

import (
	"github.com/bpremika/go-toDoList/database"
	"github.com/bpremika/go-toDoList/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ReorderTaskSchema struct {
	NewPosition int `json:"new_position"`
}

func ReorderTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := models.Task{}
	result := database.DB.Db.Select("position", "list_id").Where("id = ?", id).Find(&task)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
	body := ReorderTaskSchema{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if int(task.Position) > body.NewPosition {
		res := database.DB.Db.Model(&task).Where("list_id = ? AND position >= ? AND position < ?", task.ListID, body.NewPosition, task.Position).UpdateColumn("position", gorm.Expr("position + ?", 1))
		if res.Error != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": res.Error})
		}
		database.DB.Db.Model(&task).Where("id = ?", id).Update("position", body.NewPosition)
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "task reordered"})
	}
	if int(task.Position) < body.NewPosition {
		res := database.DB.Db.Model(&task).Where("list_id = ? AND position > ? AND position <= ?", task.ListID, task.Position, body.NewPosition).UpdateColumn("position", gorm.Expr("position - ?", 1))
		if res.Error != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": res.Error})
		}
		database.DB.Db.Model(&task).Where("id = ?", id).Update("position", body.NewPosition)
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "task reordered"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "task reordered"})
}
