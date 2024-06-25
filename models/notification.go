package models

import "time"

type Notification struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
