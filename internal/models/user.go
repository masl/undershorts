package models

import "time"

type User struct {
	Id           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
