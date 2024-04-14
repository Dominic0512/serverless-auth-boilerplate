package database

import (
	"github.com/Dominic0512/serverless-auth-boilerplate/ent"
	"github.com/google/wire"
)

type Tx *ent.Tx

var ProviderSet = wire.NewSet(
	NewPSQLDatabase,
)
