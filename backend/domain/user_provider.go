package domain

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/ent"
	"github.com/Dominic0512/serverless-auth-boilerplate/ent/userprovider"
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
	Create(userProvider UserProviderEntity) (*UserProviderEntity, error)
}
