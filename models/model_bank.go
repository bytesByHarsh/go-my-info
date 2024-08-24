package models

import (
	"database/sql"
	"time"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/google/uuid"
)

type CreateBankReq struct {
	Name            string            `json:"name"`
	ContactPhone    string            `json:"contact_phone"`
	ContactEmail    string            `json:"contact_email"`
	Address         string            `json:"address"`
	BankType        database.BankType `json:"bank_type"`
	EstablishedYear int32             `json:"established_year"`
}

type Bank struct {
	ID              uuid.UUID         `json:"id"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	Name            string            `json:"name"`
	ContactPhone    sql.NullString    `json:"contact_phone"`
	ContactEmail    string            `json:"contact_email"`
	Address         sql.NullString    `json:"address"`
	Type            database.BankType `json:"bank_type"`
	EstablishedYear int32             `json:"established_year"`
}

func ConvBankToBank(dbBank database.Bank) Bank {
	return Bank{
		ID:              dbBank.ID,
		CreatedAt:       dbBank.CreatedAt,
		UpdatedAt:       dbBank.UpdatedAt,
		Name:            dbBank.Name,
		ContactPhone:    dbBank.ContactPhone,
		ContactEmail:    dbBank.ContactEmail,
		Address:         dbBank.Address,
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
