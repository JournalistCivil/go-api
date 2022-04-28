package goapi

import internal "github.com/anon/go-api/internal/base"

type User struct {
	internal.BaseModel
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserService interface {
	FindById(id string) (*User, error)
	// All() ([]*User, error)
	// Create(u *User) error
	// Delete(id string) error
}
