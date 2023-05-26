package handlers

import (
	"fmt"

	"github.com/bpremika/go-toDoList/database"
	"github.com/bpremika/go-toDoList/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetLists(c *fiber.Ctx) error {
	lists := []models.List{}
	results := database.DB.Db.Debug().Order("position").Preload("Tasks").Find(&lists)
	if results.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": results.Error})
	}
	return c.Status(200).JSON(lists)
}

func CreateList(c *fiber.Ctx) error {
	newList := new(models.List)
	if err := c.BodyParser(newList); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	result := database.DB.Db.Create(&newList)
	fmt.Println(result.Error)
	return c.Status(200).JSON(newList)
}

func DeleteList(c *fiber.Ctx) error {
	id := c.Params("id")
	list := models.List{}
	lists := models.List{}
	task := models.Task{}
	res := database.DB.Db.Where("list_id = ?", id).Delete(&task)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": res.Error})
	}
	res = database.DB.Db.Where("id = ?", id).Find(&list)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": res.Error})
	}
	fmt.Println(list.Position)
	res = database.DB.Db.Model(&lists).Where("position > ?", list.Position).UpdateColumn("position", gorm.Expr("position - ?", 1))
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": res.Error})
	}
	result := database.DB.Db.Delete(&list, id)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
	return c.Status(200).JSON(list)
}
