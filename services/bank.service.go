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

type BankService struct {
	BankRepository repository.IBankRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBankService(bankRepository repository.IBankRepository, DB *sql.DB, validate *validator.Validate) IBankService {
	return &BankService{
		BankRepository: bankRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *BankService) Create(ctx context.Context, request dto.CreateBankDto) mapping.BankVm {
	err := service.Validate.Struct(request)
	common.PanicIfError(err)

	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	bank := entity.Bank{
		BankCode:     &request.BankCode,
		BankName:     &request.BankName,
		UserInserted: &request.UserInserted,
		InsertedDate: time.Now(),
		LastUserId:   new(string),
		LastUpdate:   time.Now(),
	}

	bank = service.BankRepository.Save(ctx, tx, bank)

	return mapping.ToBankResponse(bank)
}

func (service *BankService) Update(ctx context.Context, request dto.UpdateBankDto, oid string) mapping.BankVm {
	err := service.Validate.Struct(request)
	common.PanicIfError(err)

	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	bank, err := service.BankRepository.FindById(ctx, tx, oid)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	bank.BankCode = &request.BankCode
	bank.BankName = &request.BankName
	bank.LastUserId = &request.LastUserId

	bank = service.BankRepository.Update(ctx, tx, bank)

	return mapping.ToBankResponse(bank)
}

func (service *BankService) Delete(ctx context.Context, oid string) {
	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	bank, err := service.BankRepository.FindById(ctx, tx, oid)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	service.BankRepository.Delete(ctx, tx, bank)
}

func (service *BankService) FindById(ctx context.Context, oid string) mapping.BankVm {
	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	bank, err := service.BankRepository.FindById(ctx, tx, oid)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	return mapping.ToBankResponse(bank)
}

func (service *BankService) Find(ctx context.Context, criteria string) []mapping.BankVm {
	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	banks := service.BankRepository.Find(ctx, tx, criteria)

	return mapping.ToBankResponses(banks)
}
