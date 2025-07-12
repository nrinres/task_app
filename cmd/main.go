package main

import (
  "github.com/gofiber/fiber/v2"
  "task_app/config"
  "task_app/internal/routes"
  )

func main() {
  config.ConnectDB()
  
  app := fiber.New()
  routes.TaskRoute(app)
  
  app.Listen(":3000")
}