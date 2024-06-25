package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rishabh/practice/models"
	"github.com/rishabh/practice/usecase"
)

type NotificationHandler struct {
	NotificationUsecase *usecase.NotificationUsecase
}

func NewNotificationHandler(router *gin.Engine, notificationUsecase usecase.NotificationUsecase) {
	handler := &NotificationHandler{
		NotificationUsecase: &notificationUsecase,
	}

	router.POST("/notifications", handler.CreateNotification)
}

func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var notification models.Notification

	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.NotificationUsecase.CreateNotification(notification.Email, notification.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "notification created"})
}

func (h *NotificationHandler) ListNotification(c *gin.Context) {
	notifications, err := h.NotificationUsecase.ListNotifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.NotificationUsecase.DeletetNotifications(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "notification deleted"})
}
