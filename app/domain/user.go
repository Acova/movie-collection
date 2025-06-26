package domain

import (
	"time"
)

type User struct {
	ID           uint
	Email        string
	Name         string
	Password     string
	RegisterDate time.Time
	DisableDate  time.Time
}
