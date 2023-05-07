package repository

import (
	"context"
	"database/sql"

	"github.com/celpostgress-api/entity"
)

type IPermissionPolicyUserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, permissionPolicyUser entity.PermissionPolicyUser) entity.PermissionPolicyUser
	Update(ctx context.Context, tx *sql.Tx, permissionPolicyUser entity.PermissionPolicyUser) entity.PermissionPolicyUser
	Delete(ctx context.Context, tx *sql.Tx, permissionPolicyUser entity.PermissionPolicyUser)
	FindById(ctx context.Context, tx *sql.Tx, oid string) (entity.PermissionPolicyUser, error)
	Find(ctx context.Context, tx *sql.Tx, criteria string) []entity.PermissionPolicyUser
}
