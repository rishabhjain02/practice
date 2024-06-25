package repository

type NotificationRepoStruct struct {
	db string
}

func NewNotificationRepo(db string) NotificationRepo {
	return nil
}

// To Do

// Creating connection to db
// Function to create entry of the notification in db
// Function to Get all the notifications from the db
// Function to Delete the notification from the db
