package dto

type FilterBankDto struct {
	BankCode string `json:"bankCode"`
	BankName string `json:"bankName"`
}

type PageBankDto struct {
	Oid       string `json:"oid"`
	BankCode  string `json:"bankCode"`
	BankName  string `json:"bankName"`
	ShortName string `json:"shortName"`
}

type CreateBankDto struct {
	BankCode     string `json:"bankCode"  binding:"required"`
	BankName     string `json:"bankName"  binding:"required"`
	UserInserted string `json:"userInserted"`
}

type UpdateBankDto struct {
	BankCode   string `json:"bankCode" binding:"required"`
	BankName   string `json:"bankName" binding:"required"`
	LastUserId string `json:"lastUserId"`
}
