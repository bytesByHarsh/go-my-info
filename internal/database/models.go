// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BankAccountType string

const (
	BankAccountTypeSavings  BankAccountType = "savings"
	BankAccountTypeChecking BankAccountType = "checking"
)

func (e *BankAccountType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BankAccountType(s)
	case string:
		*e = BankAccountType(s)
	default:
		return fmt.Errorf("unsupported scan type for BankAccountType: %T", src)
	}
	return nil
}

type NullBankAccountType struct {
	BankAccountType BankAccountType
	Valid           bool // Valid is true if BankAccountType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBankAccountType) Scan(value interface{}) error {
	if value == nil {
		ns.BankAccountType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BankAccountType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBankAccountType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BankAccountType), nil
}

type BankType string

const (
	BankTypeCentral      BankType = "central"
	BankTypeCooperative  BankType = "cooperative"
	BankTypeCommercial   BankType = "commercial"
	BankTypeRegional     BankType = "regional"
	BankTypeLocal        BankType = "local"
	BankTypeSpecialized  BankType = "specialized"
	BankTypeSmallFinance BankType = "small_finance"
	BankTypePayments     BankType = "payments"
)

func (e *BankType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BankType(s)
	case string:
		*e = BankType(s)
	default:
		return fmt.Errorf("unsupported scan type for BankType: %T", src)
	}
	return nil
}

type NullBankType struct {
	BankType BankType
	Valid    bool // Valid is true if BankType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBankType) Scan(value interface{}) error {
	if value == nil {
		ns.BankType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BankType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBankType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BankType), nil
}

type CardType string

const (
	CardTypeCredit CardType = "credit"
	CardTypeDebit  CardType = "debit"
)

func (e *CardType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CardType(s)
	case string:
		*e = CardType(s)
	default:
		return fmt.Errorf("unsupported scan type for CardType: %T", src)
	}
	return nil
}

type NullCardType struct {
	CardType CardType
	Valid    bool // Valid is true if CardType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCardType) Scan(value interface{}) error {
	if value == nil {
		ns.CardType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CardType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCardType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CardType), nil
}

type Bank struct {
	ID              uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
	IsDeleted       bool
	Name            string
	ContactPhone    sql.NullString
	ContactEmail    string
	Address         sql.NullString
	Type            BankType
	EstablishedYear int32
}

type BankAccount struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
	IsDeleted     bool
	IsActive      bool
	BankID        uuid.UUID
	UserID        uuid.UUID
	Name          string
	AccountNumber string
	AccountType   BankAccountType
	Balance       string
	Currency      string
}

type Card struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
	IsDeleted     bool
	IsActive      bool
	UserID        uuid.UUID
	BankID        uuid.UUID
	BankAccountID uuid.NullUUID
	Name          string
	Nickname      string
	Number        string
	Type          CardType
	ExpMonth      int32
	ExpYear       int32
	Cvv           string
	TotalLimit    string
	BillDate      int32
}

type User struct {
	ID             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime
	IsDeleted      bool
	Name           string
	PhoneNum       string
	Email          string
	Username       string
	ProfileImg     string
	Role           int32
	HashedPassword string
	IsActive       bool
}
