package core

import "context"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type UserRepository interface {
	Save(ctx context.Context, user User) error
	GetByUsername(ctx context.Context, username string) (*User, error)
}
