package models

import "time"

type Todo struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	HasCompleted bool      `json:"hasCompleted"`
	CreatedAt    time.Time `json:"createdAt"`
}
