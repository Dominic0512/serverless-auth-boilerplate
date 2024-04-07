package helper

import (
	"context"

	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
)

type TxHelper interface {
	WithTx(ctx context.Context, fn func(tx database.Tx) error) error
}
