package repository

import "github.com/rishabh/practice/models"

type NotificationRepo interface {
	Create(Notification *models.Notification) error
	GetAll() ([]models.Notification, error)
	Delete(id int) error
}
