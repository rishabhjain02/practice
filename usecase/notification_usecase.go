package usecase

import (
	"sync"

	"github.com/rishabh/practice/delivery"
	"github.com/rishabh/practice/models"
	"github.com/rishabh/practice/repository"
)

type NotificationUsecase struct {
	NotificationRepo repository.NotificationRepo
}

func NewNotificationUsecase(notificationRepo repository.NotificationRepo) *NotificationUsecase {
	return &NotificationUsecase{
		NotificationRepo: notificationRepo,
	}
}

func (n *NotificationUsecase) CreateNotification(email string, content string) error {
	notification := &models.Notification{
		Email:   email,
		Content: content,
		Status:  "new",
	}

	err := n.NotificationRepo.Create(notification)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := delivery.SendEmail(notification.Email, notification.Content)
		if err != nil {
			notification.Status = "failed"
		} else {
			notification.Status = "sent"
		}

		n.NotificationRepo.Create(notification)
	}()

	wg.Wait()
	return nil
}

func (n *NotificationUsecase) ListNotifications() ([]models.Notification, error) {
	return n.NotificationRepo.GetAll()
}

func (n *NotificationUsecase) DeletetNotifications(id int) error {
	return n.NotificationRepo.Delete(id)
}
