package domain

import (
	"context"

	"github.com/Dominic0512/serverless-auth-boilerplate/ent"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/userprovider"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
)

const (
	UserProviderNamePrimary  = userprovider.NamePrimary
	UserProviderNameFacebook = userprovider.NameFacebook
	UserProviderNameGoogle   = userprovider.NameGoogle
)

type UserProviderEntity = ent.UserProvider

type UserProviderClient = ent.UserProviderClient

type UserProvider struct {
	Name string
}

type UserProviderRepository interface {
	Create(ctx context.Context, tx database.Tx, userProvider UserProviderEntity) (*UserProviderEntity, error)
}
