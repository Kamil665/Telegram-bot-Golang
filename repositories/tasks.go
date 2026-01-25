package repositories

import (
	"telegram-NewBot/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateTask(userID int64, text string) (*models.Task, error) {
	task := &models.Task{
		ChatID: userID,
		Task:   text,
	}

	if err := DB.Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}
