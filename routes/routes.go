package routes

import (
	"github.com/bpremika/go-toDoList/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupListsRoutes(router fiber.Router) {
	router.Get("/", handlers.GetLists)
	router.Post("/", handlers.CreateList)
	router.Patch("/:id", handlers.UpdateList)
	router.Delete("/:id", handlers.DeleteList)
	router.Patch("/reorder/:id", handlers.ReorderList)
}

func SetupTasksRoutes(router fiber.Router) {
	router.Get("/", handlers.GetTasks)
	router.Post("/", handlers.CreateTask)
	router.Patch("/:id", handlers.UpdateTask)
	router.Delete("/:id", handlers.DeleteTask)
	router.Patch("/reorder/:id", handlers.ReorderTask)
	router.Patch("/movetonewlist/:id", handlers.MoveTasktoAnotherList)
}
