package entity

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todo struct {
	ID        uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
