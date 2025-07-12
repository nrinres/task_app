package routes


import (
  "github.com/gofiber/fiber/v2"
  "task_app/internal/task/controller"
)

func TaskRoute(app *fiber.App) {
task := app.Group("/api/v1/task")
task.Post("/", controller.CreateTask)
task.Get("/", controller.GetTasks)
task.Put("/:id", controller.UpdateTask)
task.Delete("/:id", controller.DeleteTask)
}