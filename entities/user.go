package entities

import (
	"time"
)

type User struct {
	Id        uint
	Username  *Username
	Password  *Password
	CreatedAt time.Time
}
