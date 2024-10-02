package models

import (
	"time"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/google/uuid"
)

type AddBankAccountReq struct {
	BankID        string `json:"bank_id"  validate:"required"`
	AccountNumber string `json:"account_number"  validate:"required"`
	AccountType   string `json:"account_type"  validate:"required"`
	Name          string `json:"name"  validate:"required"`
	Balance       string `json:"balance"  validate:"required"`
	Currency      string `json:"currency"  validate:"required"`
}

type UpdateBankAccountReq struct {
	AccountNumber string `json:"account_number"  validate:"required"`
	AccountType   string `json:"account_type"  validate:"required"`
	Name          string `json:"name"  validate:"required"`
	Balance       string `json:"balance"  validate:"required"`
	Currency      string `json:"currency"  validate:"required"`
}

type BankAccount struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	AccountNumber string    `json:"account_number"`
	AccountType   string    `json:"account_type"`
	Balance       string    `json:"balance"`
	Currency      string    `json:"currency"`
}

func ConvAccountToAccount(dbBankAccount database.BankAccount) BankAccount {
	return BankAccount{
		ID:            dbBankAccount.ID,
		CreatedAt:     dbBankAccount.CreatedAt,
		UpdatedAt:     dbBankAccount.UpdatedAt,
		Name:          dbBankAccount.Name,
		AccountNumber: dbBankAccount.AccountNumber,
		AccountType:   string(dbBankAccount.AccountType),
		Balance:       dbBankAccount.Balance,
		Currency:      dbBankAccount.Currency,
	}
}

func CreateAccountListResp(dbBankAccountList []database.BankAccount) []BankAccount {
	bankAccountList := []BankAccount{}
	for _, dbBank := range dbBankAccountList {
		bankAccountList = append(bankAccountList, ConvAccountToAccount(dbBank))
	}
	return bankAccountList
}
