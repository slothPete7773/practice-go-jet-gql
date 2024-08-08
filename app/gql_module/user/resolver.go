package user

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

type UserResolver struct{}

func (UserResolver) CreateUser(ctx context.Context, input CreateUserInput) (*User, error)

func (UserResolver) UpdateUser(ctx context.Context, input UpdateUserInput) (*User, error)

func (UserResolver) DeleteUser(ctx context.Context, input DeleteUserInput) (*User, error)

func (UserResolver) GetUser(ctx context.Context, input GetUserInput) (*User, error)

func (UserResolver) GetAllUsers(ctx context.Context) (*[]User, error)

type User struct {
	ID       graphql.ID `json:"id"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Name     *string    `json:"name"`
	Role     string     `json:"role"`
}

type GetUserInput struct {
	ID graphql.ID `json:"id"`
}

type CreateUserInput struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Name     *string `json:"name"`
	Role     string  `json:"role"`
}

type UpdateUserInput struct {
	ID       graphql.ID `json:"id"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Name     *string    `json:"name"`
	Role     string     `json:"role"`
}

type DeleteUserInput struct {
	ID graphql.ID `json:"id"`
}
