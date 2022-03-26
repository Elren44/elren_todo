package model

import "time"

type User struct {
	UUID              string    `json:"uuid"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"created_at"`
}

type UserDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}
