package models

import (
	"time"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/google/uuid"
)

type CreateBankReq struct {
	Name            string            `json:"name" validate:"required"`
	ContactPhone    string            `json:"contact_phone"`
	ContactEmail    string            `json:"contact_email" validate:"required"`
	Address         string            `json:"address" validate:"required"`
	BankType        database.BankType `json:"bank_type" validate:"required"`
	EstablishedYear int32             `json:"established_year" validate:"required"`
}

type UpdateBankReq struct {
	Name         string            `json:"name" validate:"required"`
	ContactPhone string            `json:"contact_phone"`
	ContactEmail string            `json:"contact_email" validate:"required"`
	Address      string            `json:"address" validate:"required"`
	BankType     database.BankType `json:"bank_type" validate:"required"`
}

type Bank struct {
	ID              uuid.UUID         `json:"id"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	Name            string            `json:"name"`
	ContactPhone    *string           `json:"contact_phone"`
	ContactEmail    string            `json:"contact_email"`
	Address         *string           `json:"address"`
	Type            database.BankType `json:"bank_type"`
	EstablishedYear int32             `json:"established_year"`
}

func ConvBankToBank(dbBank database.Bank) Bank {
	var phone_num *string
	var address *string
	if !dbBank.ContactPhone.Valid {
		dbBank.ContactPhone.String = "-"
		dbBank.ContactPhone.Valid = true
	}
	if !dbBank.Address.Valid {
		dbBank.Address.String = "-"
		dbBank.Address.Valid = true
	}
	phone_num = &dbBank.ContactPhone.String
	address = &dbBank.Address.String
	return Bank{
		ID:              dbBank.ID,
		CreatedAt:       dbBank.CreatedAt,
		UpdatedAt:       dbBank.UpdatedAt,
		Name:            dbBank.Name,
		ContactPhone:    phone_num,
		ContactEmail:    dbBank.ContactEmail,
		Address:         address,
		Type:            dbBank.Type,
		EstablishedYear: dbBank.EstablishedYear,
	}
}

func CreateBankListResp(dbBankList []database.Bank) []Bank {
	bankList := []Bank{}
	for _, dbBank := range dbBankList {
		bankList = append(bankList, ConvBankToBank(dbBank))
	}
	return bankList
}
