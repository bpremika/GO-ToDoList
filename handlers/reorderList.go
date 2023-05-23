package handlers

import (
	"github.com/bpremika/go-toDoList/database"
	"github.com/bpremika/go-toDoList/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ReorderListSchema struct {
	NewPosition int `json:"new_position"`
}

func ReorderList(c *fiber.Ctx) error {
	id := c.Params("id")
	list := models.List{}
	result := database.DB.Db.Select("position").Where("id = ?", id).Find(&list)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
	body := ReorderListSchema{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if int(list.Position) > body.NewPosition {
		res := database.DB.Db.Model(&list).Where("position >= ? AND position < ?", body.NewPosition, list.Position).UpdateColumn("position", gorm.Expr("position + ?", 1))
		if res.Error != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": res.Error})
		}
		database.DB.Db.Model(&list).Where("id = ?", id).Update("position", body.NewPosition)
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "list reordered"})
	}
	if int(list.Position) < body.NewPosition {
		res := database.DB.Db.Model(&list).Where("position > ? AND position <= ?", list.Position, body.NewPosition).UpdateColumn("position", gorm.Expr("position - ?", 1))
		if res.Error != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": res.Error})
		}
		database.DB.Db.Model(&list).Where("id = ?", id).Update("position", body.NewPosition)
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "list reordered"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "list reordered"})
}
