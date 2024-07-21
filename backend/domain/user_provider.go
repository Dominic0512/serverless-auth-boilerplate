package domain

import (
	"context"

	"github.com/Dominic0512/serverless-auth-boilerplate/ent"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/userprovider"
	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
)

type UserProviderName = userprovider.Name

const (
	UserProviderNamePrimary  UserProviderName = userprovider.NamePrimary
	UserProviderNameAuth0    UserProviderName = userprovider.NameAuth0
	UserProviderNameFacebook UserProviderName = userprovider.NameFacebook
	UserProviderNameGoogle   UserProviderName = userprovider.NameGoogle
)

type UserProviderEntity = ent.UserProvider

type UserProviderClient = ent.UserProviderClient

type UserProviderRepository interface {
	Create(ctx context.Context, tx database.Tx, userProvider UserProviderEntity) (*UserProviderEntity, error)
}
