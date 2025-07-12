package repository

import (
  "task_app/config"
  "task_app/internal/model"
  )

func CreateTask(task *model.Task) error {
  return config.DB.Create(task).Error
}

func UpdateTask(task *model.Task) error {
  return config.DB.Save(task).Error
}

func DeleteTask(id uint) error {
	return config.DB.Delete(&model.Task{}, id).Error
}

func GetTaskByID(id uint) (*model.Task, error) {
  var task model.Task
  if err := config.DB.First(&task, "id = ?", id).Error; err != nil {
    return nil, err
  }
  return &task, nil
}

func GetAllTasks()([]model.Task, error) {
  var tasks []model.Task
  if err := config.DB.Find(&tasks).Error; err != nil {
    return nil, err
  }
  return tasks, nil
}