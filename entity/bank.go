package entity

import (
	"time"
)

type Bank struct {
	Oid          string    `json:"oid"`
	BankCode     *string   `json:"bankCode"`
	BankName     *string   `json:"bankName"`
	UserInserted *string   `json:"userInserted"`
	InsertedDate time.Time `json:"insertedDate"`
	LastUserId   *string   `json:"lastUserId"`
	LastUpdate   time.Time `json:"lastUpdate"`
}
