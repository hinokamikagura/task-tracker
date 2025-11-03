package schemas

import "time"

type Task struct {
	Id          string
	Title       string
	Description string
	Status      string
}

type TaskResponse struct {
	Id          string    `json:id`
	Title       string    `json:title`
	Description string    `json:description`
	Status      string    `json:status`
	CreatedAt   time.Time `json:createdAt`
	UpdatedAt   time.Time `json:updatedAt`
}
