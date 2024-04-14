package helper

import (
	"context"
	"fmt"

	"github.com/Dominic0512/serverless-auth-boilerplate/infra/database"
)

type EntTxHelper struct {
	db *database.PSQLDatabase
}

func (th *EntTxHelper) WithTx(ctx context.Context, fn func(tx database.Tx) error) error {
	tx, err := th.db.Client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func NewEntTxHelper(db *database.PSQLDatabase) *EntTxHelper {
	return &EntTxHelper{
		db,
	}
}
