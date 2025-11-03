package schemas

import "time"

// type Task struct {
// 	Id          string
// 	Title       string
// 	Description string
// 	Status      string
// }

type Task struct {
	Id          uint      `json:"Id"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	Status      string    `json:"Status"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

type TaskOutPut struct {
	Id          uint   `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Status      string `json:"Status"`
	CreatedAt   string `json:"CreatedAt"`
	UpdatedAt   string `json:"UpdatedAt"`
}
