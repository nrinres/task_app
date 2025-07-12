package config

import (
  "log"
  "os"
  "github.com/joho/godotenv"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  
  "task_app/internal/model"
  )

var DB *gorm.DB
func ConnectDB() {
  if err := godotenv.Load(); err != nil {
    log.Println("No .env file found")
  }
  
  dsn := os.Getenv("DATABASE_URL")
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed connect to database")
  }
  DB = db
  
  if err := DB.AutoMigrate(&model.Task{}); err != nil {
    panic("failed to auto migrate: " + err.Error())
  }
}