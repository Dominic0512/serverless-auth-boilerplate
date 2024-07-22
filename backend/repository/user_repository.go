package repository

import (
	"context"

	"github.com/Dominic0512/serverless-auth-boilerplate/domain"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/user"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
	"github.com/google/uuid"
)

type UserRepository struct {
	User *domain.UserClient
}

func (ur UserRepository) Find(ctx context.Context) ([]*domain.UserEntity, error) {
	user, err := ur.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur UserRepository) FindOne(ctx context.Context, id uuid.UUID) (*domain.UserEntity, error) {
	user, err := ur.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) FindOneByEmail(ctx context.Context, email string) (*domain.UserEntity, error) {
	user, err := ur.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) Create(ctx context.Context, tx database.Tx, properties domain.UserEntity) (*domain.UserEntity, error) {
	repo := ur.User

	if tx != nil {
		repo = tx.User
	}

	mutate := repo.Create().
		SetEmail(properties.Email).
		SetName(properties.Name)

	if properties.Password != nil {
		mutate.SetPassword(*properties.Password)
	}

	user, err := mutate.Save(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) Update(ctx context.Context, id uuid.UUID, properties domain.UserEntity) (*domain.UserEntity, error) {
	user, err := ur.User.UpdateOneID(id).
		SetName(properties.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := ur.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *database.PSQLDatabase) *UserRepository {
	return &UserRepository{
		User: db.Client.User,
	}
}
