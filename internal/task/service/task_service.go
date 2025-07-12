package service

import (
  "task_app/internal/model"
  "task_app/internal/task/repository"
  )

func CreateTask(title string, completed bool) (*model.Task, error) {
  task := &model.Task{
    Title:      title,
    Completed:   completed,
  }
  if err := repository.CreateTask(task); err != nil {
    return nil, err
  }
  return task, nil
}

func UpdateTask(id uint, title string, completed bool) error {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		return err
	}
	task.Title = title
	task.Completed = completed
	return repository.UpdateTask(task)
}

func GetAllTasks()([]model.Task, error) {
  return repository.GetAllTasks()
}

func DeleteTask(id uint) error {
  return repository.DeleteTask(id)
}