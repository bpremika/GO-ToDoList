package main

import (
	"log"

	"github.com/bpremika/go-toDoList/database"
	"github.com/bpremika/go-toDoList/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	api := app.Group("/api")

	lists := api.Group("/lists")
	tasks := api.Group("/tasks")

	routes.SetupListsRoutes(lists)
	routes.SetupTasksRoutes(tasks)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, !")
	})

	log.Fatal(app.Listen(":3000"))
}
