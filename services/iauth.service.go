package services

import (
	"context"

	"github.com/celpostgress-api/dto"
	"github.com/celpostgress-api/entity"
	"github.com/celpostgress-api/mapping"
)

type IAuthService interface {
	Login(ctx context.Context, criteria string) entity.PermissionPolicyUser
	Register(ctx context.Context, request dto.PayloadRegister) mapping.PermissionPolicyUserVm
}
