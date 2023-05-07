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
	"golang.org/x/crypto/bcrypt"
)

type PermissionPolicyUserService struct {
	PermissionPolicyUserRepository repository.IPermissionPolicyUserRepository
	DB                             *sql.DB
	Validate                       *validator.Validate
}

func NewPermissionPolicyUserService(permissionPolicyUserRepository repository.IPermissionPolicyUserRepository, DB *sql.DB, validate *validator.Validate) IPermissionPolicyUserService {
	return &PermissionPolicyUserService{
		PermissionPolicyUserRepository: permissionPolicyUserRepository,
		DB:                             DB,
		Validate:                       validate,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (service *PermissionPolicyUserService) Create(ctx context.Context, request dto.CreatePermissionPolicyUserDto) mapping.PermissionPolicyUserVm {
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
		UserInserted:        &request.UserInserted,
		InsertedDate:        time.Now(),
		LastUserId:          new(string),
		LastUpdate:          time.Now(),
	}

	permissionPolicyUser = service.PermissionPolicyUserRepository.Save(ctx, tx, permissionPolicyUser)

	return mapping.ToPermissionPolicyUserResponse(permissionPolicyUser)
}

func (service *PermissionPolicyUserService) Update(ctx context.Context, request dto.UpdatePermissionPolicyUserDto, oid string) mapping.PermissionPolicyUserVm {
	err := service.Validate.Struct(request)
	common.PanicIfError(err)

	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	permissionPolicyUser, err := service.PermissionPolicyUserRepository.FindById(ctx, tx, oid)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}
	permissionPolicyUser.EmailName = request.EmailName
	permissionPolicyUser.Password = request.Password
	permissionPolicyUser.LevelUser = request.LevelUser
	permissionPolicyUser.AndroidToken = &request.AndroidToken
	permissionPolicyUser.LastUserId = &request.LastUserId

	permissionPolicyUser = service.PermissionPolicyUserRepository.Update(ctx, tx, permissionPolicyUser)

	return mapping.ToPermissionPolicyUserResponse(permissionPolicyUser)
}

func (service *PermissionPolicyUserService) Delete(ctx context.Context, oid string) {
	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	permissionPolicyUser, err := service.PermissionPolicyUserRepository.FindById(ctx, tx, oid)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	service.PermissionPolicyUserRepository.Delete(ctx, tx, permissionPolicyUser)
}

func (service *PermissionPolicyUserService) FindById(ctx context.Context, oid string) mapping.PermissionPolicyUserVm {
	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	permissionPolicyUser, err := service.PermissionPolicyUserRepository.FindById(ctx, tx, oid)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	return mapping.ToPermissionPolicyUserResponse(permissionPolicyUser)
}

func (service *PermissionPolicyUserService) Find(ctx context.Context, criteria string) []mapping.PermissionPolicyUserVm {
	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	permissionPolicyUsers := service.PermissionPolicyUserRepository.Find(ctx, tx, criteria)

	return mapping.ToPermissionPolicyUserResponses(permissionPolicyUsers)
}
