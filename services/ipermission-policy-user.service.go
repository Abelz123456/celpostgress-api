package services

import (
	"context"

	"github.com/celpostgress-api/dto"
	"github.com/celpostgress-api/mapping"
)

type IPermissionPolicyUserService interface {
	Create(ctx context.Context, request dto.CreatePermissionPolicyUserDto) mapping.PermissionPolicyUserVm
	Update(ctx context.Context, request dto.UpdatePermissionPolicyUserDto, oid string) mapping.PermissionPolicyUserVm
	Delete(ctx context.Context, oid string)
	FindById(ctx context.Context, oid string) mapping.PermissionPolicyUserVm
	Find(ctx context.Context, criteria string) []mapping.PermissionPolicyUserVm
}
