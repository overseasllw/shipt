package models

import "time"

type User struct {
	Id        int64     `json:"id"`
	Username  *string   `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Firstname *string   `json:"firstname"`
	Lastname  *string   `json:"lastname"`
	CreatedAt time.Time `json:"created_at"`
}
