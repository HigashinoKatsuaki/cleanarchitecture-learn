package model

import "time"

type Todo struct {
	ID        int64     `json:"id" db:"id"`
	Status    string    `json:"status" db:"status"`
	Task      string    `json:"task" db:"task"`
	LimitDate time.Time `json:"limit_date" db:"limit_date"`
}
