package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/Dominic0512/serverless-auth-boilerplate/ent"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/user"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
)

type UserRepository struct {
	db *database.Database
}

func (ur UserRepository) GetAllUsers(name string) ([]*ent.User, error) {
	u, err := ur.db.Client.User.Query().Where(user.NameContains(name)).All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func (ur UserRepository) GetUser() (*ent.User, error) {
	return nil, nil
}

func (ur UserRepository) CreateUser(user ent.User) (*ent.User, error) {
	u, err := ur.db.Client.User.Create().
		SetEmail(user.Email).
		SetPassword(*user.Password).
		SetPasswordSalt(*user.PasswordSalt).
		SetName(user.Name).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed mutating user: %w", err)
	}

	return u, nil
}

func (ur UserRepository) UpdateUser() (*ent.User, error) {
	return nil, nil
}

func (ur UserRepository) DeleteUser() (*ent.User, error) {
	return nil, nil
}

func NewUserRepository(db *database.Database) UserRepository {
	return UserRepository{db}
}
