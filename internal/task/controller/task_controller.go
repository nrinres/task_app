package controller


import (
  "strconv"
  "github.com/gofiber/fiber/v2"
  "task_app/internal/task/service"
  )

type CreateRequest struct {
  Title     string `json:"title"`
  Completed  bool `json:"completed"`
}

func CreateTask(c *fiber.Ctx) error {
  var body CreateRequest
  if err := c.BodyParser(&body); err != nil {
    return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
  }
  
  task, err := service.CreateTask(body.Title, body.Completed)
  if err != nil {
    return c.Status(500).JSON(fiber.Map{"error": err.Error()})
  }
  return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
  id := c.Params("id")
  
  idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid task ID"})
	}
	type Body struct {
	  Title string `json:"title"`
	  Completed bool
	  
	}
	var body Body
	if err := c.BodyParser(&body); err != nil {
	  return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := service.UpdateTask(uint(idUint), body.Title, body.Completed); err != nil {
	  return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Task updated"})
}

func GetTasks(c *fiber.Ctx) error {
  tasks, err := service.GetAllTasks()
  if err != nil {
    return c.Status(500).JSON(fiber.Map{"error": err.Error()})
  }
  return c.JSON(tasks)
}

func DeleteTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	if err := service.DeleteTask(uint(idUint)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Task deleted successfully"})
}
