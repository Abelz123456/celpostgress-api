package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/celpostgress-api/common"
	"github.com/celpostgress-api/entity"

	"github.com/feiin/ploto"
	"github.com/google/uuid"
)

type BankRepository struct {
}

func NewBankRepository() IBankRepository {
	return &BankRepository{}
}

func (repository *BankRepository) Save(ctx context.Context, tx *sql.Tx, bank entity.Bank) entity.Bank {

	id := uuid.New()
	SQL := `INSERT INTO public.bank(oid, bankcode, bankname,insertedDate,userInserted) VALUES($1,$2,$3,$4,$5)`

	_, err := tx.ExecContext(ctx, SQL, id, bank.BankCode, bank.BankName,
		bank.InsertedDate, bank.UserInserted)
	common.PanicIfError(err)
	bank.Oid = id.String()
	return bank
}

func (repository *BankRepository) Update(ctx context.Context, tx *sql.Tx, bank entity.Bank) entity.Bank {

	SQL := `update bank set "lastUserId"=$1, "lastUpdate"=$2 ,"bankCode" = $3,"bankName" = $4 where "oid" = $5`
	_, err := tx.ExecContext(ctx, SQL, bank.LastUserId, bank.LastUpdate, bank.BankCode, bank.BankName, bank.Oid)
	common.PanicIfError(err)

	return bank
}

func (repository *BankRepository) Delete(ctx context.Context, tx *sql.Tx, bank entity.Bank) {
	SQL := "delete from bank where oid = $1"
	_, err := tx.ExecContext(ctx, SQL, bank.Oid)
	common.PanicIfError(err)
}

func (repository *BankRepository) FindById(ctx context.Context, tx *sql.Tx, oid string) (entity.Bank, error) {
	SQL := `select  "userInserted", "insertedDate", "lastUserId", "oid", "bankCode","bankName" from bank where "oid" = $1`
	rows, err := tx.QueryContext(ctx, SQL, oid)
	common.PanicIfError(err)
	defer rows.Close()

	bank := entity.Bank{}
	if rows.Next() {
		err := ploto.Scan(rows, &bank)
		common.PanicIfError(err)
		return bank, nil
	} else {
		return bank, errors.New("bank is not found")
	}
}

func (repository *BankRepository) Find(ctx context.Context, tx *sql.Tx, criteria string) []entity.Bank {
	SQL := `select "userInserted", "insertedDate", "lastUserId","oid","bankCode","bankName" from public.bank ` + criteria
	rows, err := tx.QueryContext(ctx, SQL)
	common.PanicIfError(err)
	defer rows.Close()

	var banks []entity.Bank
	for rows.Next() {
		bank := entity.Bank{}
		err := ploto.Scan(rows, &bank)
		common.PanicIfError(err)
		banks = append(banks, bank)
	}
	return banks
}
