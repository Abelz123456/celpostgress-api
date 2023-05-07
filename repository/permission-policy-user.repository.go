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

type PermissionPolicyUserRepository struct {
}

func NewPermissionPolicyUserRepository() IPermissionPolicyUserRepository {
	return &PermissionPolicyUserRepository{}
}

func (repository *PermissionPolicyUserRepository) Save(ctx context.Context, tx *sql.Tx, permissionPolicyUser entity.PermissionPolicyUser) entity.PermissionPolicyUser {

	id := uuid.New()
	SQL := `insert into "permissionPolicyUser" ("userInserted", "insertedDate","oid","emailName","password") values ($1,$2,$3,$4,$5)`
	_, err := tx.ExecContext(ctx, SQL, permissionPolicyUser.UserInserted,
		permissionPolicyUser.InsertedDate, id,
		permissionPolicyUser.EmailName, permissionPolicyUser.Password)
	common.PanicIfError(err)
	permissionPolicyUser.Oid = id.String()
	return permissionPolicyUser
}

func (repository *PermissionPolicyUserRepository) Update(ctx context.Context, tx *sql.Tx, permissionPolicyUser entity.PermissionPolicyUser) entity.PermissionPolicyUser {

	SQL := "update permissionPolicyUser set  optimisticLockField=?, gCRecord=?, deleted=?, lastUserId=?, lastUpdate=? where oid = ?"
	_, err := tx.ExecContext(ctx, SQL, permissionPolicyUser.Description, permissionPolicyUser.OptimisticLockField, permissionPolicyUser.GCRecord, permissionPolicyUser.Deleted, permissionPolicyUser.LastUserId, permissionPolicyUser.LastUpdate, permissionPolicyUser.Oid)
	common.PanicIfError(err)

	return permissionPolicyUser
}

func (repository *PermissionPolicyUserRepository) Delete(ctx context.Context, tx *sql.Tx, permissionPolicyUser entity.PermissionPolicyUser) {
	SQL := "delete from permissionPolicyUser where oid = ?"
	_, err := tx.ExecContext(ctx, SQL, permissionPolicyUser.Oid)
	common.PanicIfError(err)
}

func (repository *PermissionPolicyUserRepository) FindById(ctx context.Context, tx *sql.Tx, oid string) (entity.PermissionPolicyUser, error) {
	SQL := "select description, optimisticLockField, gCRecord, deleted, userInserted, insertedDate, lastUserId, lastUpdate, oid, companyName,address from permissionPolicyUser where oid = ?"
	rows, err := tx.QueryContext(ctx, SQL, oid)
	common.PanicIfError(err)
	defer rows.Close()

	permissionPolicyUser := entity.PermissionPolicyUser{}
	if rows.Next() {
		err := ploto.Scan(rows, &permissionPolicyUser)
		common.PanicIfError(err)
		return permissionPolicyUser, nil
	} else {
		return permissionPolicyUser, errors.New("permissionPolicyUser is not found")
	}
}

func (repository *PermissionPolicyUserRepository) Find(ctx context.Context, tx *sql.Tx, criteria string) []entity.PermissionPolicyUser {
	SQL := `select  "userInserted", "insertedDate", "lastUserId", 
		  "oid","emailName","password" from "permissionPolicyUser" ` + criteria
	rows, err := tx.QueryContext(ctx, SQL)
	common.PanicIfError(err)
	defer rows.Close()

	var permissionPolicyUsers []entity.PermissionPolicyUser
	for rows.Next() {
		permissionPolicyUser := entity.PermissionPolicyUser{}
		err := ploto.Scan(rows, &permissionPolicyUser)
		common.PanicIfError(err)
		permissionPolicyUsers = append(permissionPolicyUsers, permissionPolicyUser)
	}
	return permissionPolicyUsers
}
