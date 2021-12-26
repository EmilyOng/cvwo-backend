package models

import (
	"main/db"
	"time"
)

type Task struct {
	CommonModel
	Name        string     `gorm:"not null" json:"name"`
	Description string     `gorm:"default:''" json:"description"`
	DueAt       *time.Time `json:"dueAt"`
	Tags        *[]Tag     `gorm:"many2many:task_tags" json:"tags"`
	UserID      uint8      `json:"-"`                       // Owner of the task
	BoardID     uint8      `json:"-"`                       // Board that the task belongs to
	StateID     uint8      `gorm:"not null" json:"stateId"` // State that the task is at
}

func (task *Task) Create() error {
	result := db.DB.Create(task)
	return result.Error
}

func (task *Task) Get() error {
	result := db.DB.Model(task).Preload("Tags").Find(task)
	return result.Error
}

func (task *Task) Update() error {
	if task.Tags != nil {
		err := db.DB.Model(task).Association("Tags").Replace(task.Tags)
		if err != nil {
			return err
		}
	}
	result := db.DB.Model(task).Preload("Tags").Save(task)
	return result.Error
}

func (task *Task) Delete() error {
	err := task.Get()
	if err != nil {
		return err
	}
	err = db.DB.Debug().Model(task).Association("Tags").Delete(task.Tags)
	if err != nil {
		return err
	}
	result := db.DB.Delete(task)
	return result.Error
}
