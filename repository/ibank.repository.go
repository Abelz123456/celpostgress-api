package repository

import (
	"context"
	"database/sql"

	"github.com/celpostgress-api/entity"
)

type IBankRepository interface {
	Save(ctx context.Context, tx *sql.Tx, bank entity.Bank) entity.Bank
	Update(ctx context.Context, tx *sql.Tx, bank entity.Bank) entity.Bank
	Delete(ctx context.Context, tx *sql.Tx, bank entity.Bank)
	FindById(ctx context.Context, tx *sql.Tx, oid string) (entity.Bank, error)
	Find(ctx context.Context, tx *sql.Tx, criteria string) []entity.Bank
}
