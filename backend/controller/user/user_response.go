package user

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

type UserResponse struct {
	User User `json:"user"`
}

type UsersResponse struct {
	Users []User `json:"users"`
}
