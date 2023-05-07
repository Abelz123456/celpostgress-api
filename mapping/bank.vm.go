package mapping

import (
	"github.com/celpostgress-api/entity"
)

type BankVm struct {
	Oid      string  `json:"oid"`
	BankCode *string `json:"bankCode"`
	BankName *string `json:"bankName"`
}

func ToBankResponse(bank entity.Bank) BankVm {
	return BankVm{
		Oid:      bank.Oid,
		BankCode: bank.BankCode,
		BankName: bank.BankName,
	}
}

func ToBankResponses(banks []entity.Bank) []BankVm {
	var bankResponses []BankVm
	for _, bank := range banks {
		bankResponses = append(bankResponses, ToBankResponse(bank))
	}
	return bankResponses
}
