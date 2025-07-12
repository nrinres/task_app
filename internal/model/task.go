package model


import "time"

type Task struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}