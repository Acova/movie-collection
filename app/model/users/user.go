package users

import "time"

type User struct {
	Email        string
	Name         string
	Password     string
	RegisterDate time.Time
	DisableDate  time.Time
}
