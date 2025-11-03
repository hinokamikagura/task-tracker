package schemas

import "time"

// type Task struct {
// 	Id          string
// 	Title       string
// 	Description string
// 	Status      string
// }

type Task struct {
	Id          uint
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
