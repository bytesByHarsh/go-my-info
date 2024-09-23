package models

import (
	"time"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/google/uuid"
)

type AddCardReq struct {
	BankID          uuid.UUID `json:"bank_id"  validate:"required"`
	BankAccountID   string    `json:"account_number"`
	CardType        string    `json:"card_type"  validate:"required"`
	Name            string    `json:"name"  validate:"required"`
	Nickname        string    `json:"nickname"  validate:"required"`
	Number          string    `json:"number"  validate:"required"`
	ExpirationMonth int32     `json:"exp_month"  validate:"required"`
	ExpirationYear  int32     `json:"exp_year"  validate:"required"`
	Cvv             string    `json:"cvv"  validate:"required"`
	TotalLimit      string    `json:"total_limit"  validate:"required"`
	BillDate        int32     `json:"bill_date"  validate:"required"`
}

type UpdateCardReq struct {
	BankID          uuid.UUID `json:"bank_id"  validate:"required"`
	BankAccountID   string    `json:"account_number"`
	CardType        string    `json:"card_type"  validate:"required"`
	Name            string    `json:"name"  validate:"required"`
	Nickname        string    `json:"nickname"  validate:"required"`
	Number          string    `json:"number"  validate:"required"`
	ExpirationMonth int32     `json:"exp_month"  validate:"required"`
	ExpirationYear  int32     `json:"exp_year"  validate:"required"`
	Cvv             string    `json:"cvv"  validate:"required"`
	TotalLimit      string    `json:"total_limit"  validate:"required"`
	BillDate        int32     `json:"bill_date"  validate:"required"`
}

type Card struct {
	ID              uuid.UUID `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	BankID          string    `json:"bank_id"`
	BankAccountID   string    `json:"account_number"`
	CardType        string    `json:"card_type"`
	Name            string    `json:"name"`
	Nickname        string    `json:"nickname"`
	Number          string    `json:"number"`
	ExpirationMonth int32     `json:"exp_month"`
	ExpirationYear  int32     `json:"exp_year"`
	Cvv             string    `json:"cvv"`
	TotalLimit      string    `json:"total_limit"`
	BillDate        int32     `json:"bill_date"`
}

func ConvCardToCard(dbCard database.Card) Card {
	var account_id string
	if !dbCard.BankAccountID.Valid {
		account_id = "-"
	} else {
		account_id = dbCard.BankAccountID.UUID.String()
	}
	return Card{
		ID:              dbCard.ID,
		CreatedAt:       dbCard.CreatedAt,
		UpdatedAt:       dbCard.UpdatedAt,
		Name:            dbCard.Name,
		Nickname:        dbCard.Nickname,
		Number:          dbCard.Number,
		ExpirationMonth: dbCard.ExpMonth,
		ExpirationYear:  dbCard.ExpYear,
		Cvv:             dbCard.Cvv,
		TotalLimit:      dbCard.TotalLimit,
		BillDate:        dbCard.BillDate,
		BankID:          dbCard.BankID.String(),
		BankAccountID:   account_id,
		CardType:        string(dbCard.Type),
	}
}

func CreateCardListResp(dbCardList []database.Card) []Card {
	cardList := []Card{}
	for _, dbCard := range dbCardList {
		cardList = append(cardList, ConvCardToCard(dbCard))
	}
	return cardList
}
