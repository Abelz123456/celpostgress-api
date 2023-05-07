package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/celpostgress-api/common"
	"github.com/celpostgress-api/dto"
	"github.com/celpostgress-api/entity"
	"github.com/celpostgress-api/mapping"
	"github.com/celpostgress-api/repository"
	"github.com/go-playground/validator/v10"
)

type AuthService struct {
	AuthRepository repository.IPermissionPolicyUserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.IPermissionPolicyUserRepository, DB *sql.DB, validate *validator.Validate) IAuthService {
	return &AuthService{
		AuthRepository: authRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *AuthService) Login(ctx context.Context, criteria string) entity.PermissionPolicyUser {
	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	users := service.AuthRepository.Find(ctx, tx, criteria)

	return users[0]
}

// Register implements IAuthService
func (service *AuthService) Register(ctx context.Context, request dto.PayloadRegister) mapping.PermissionPolicyUserVm {
	err := service.Validate.Struct(request)
	common.PanicIfError(err)

	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	hash, _ := HashPassword(request.Password)

	permissionPolicyUser := entity.PermissionPolicyUser{
		EmailName:           request.EmailName,
		Password:            hash,
		Description:         new(string),
		OptimisticLockField: 0,
		GCRecord:            0,
		Deleted:             false,
		UserInserted:        new(string),
		InsertedDate:        time.Now(),
		LastUserId:          new(string),
		LastUpdate:          time.Now(),
	}

	permissionPolicyUser = service.AuthRepository.Save(ctx, tx, permissionPolicyUser)

	return mapping.ToPermissionPolicyUserResponse(permissionPolicyUser)
}
