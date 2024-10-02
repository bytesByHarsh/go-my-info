package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/google/uuid"
)

type CreateBankParams struct {
	Name            string            `json:"name"`
	ContactPhone    string            `json:"contact_phone"`
	ContactEmail    string            `json:"contact_email"`
	Address         string            `json:"address"`
	Type            database.BankType `json:"type"`
	EstablishedYear int32             `json:"established_year"`
}

func InitBankDb() {
	filename := "database/banks.json"
	// Parse the JSON file
	banks, err := ParseBanksJSON(filename)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	for _, bank := range banks {
		contactPhone := sql.NullString{}
		if bank.ContactPhone != "" {
			contactPhone.Valid = true
			contactPhone.String = bank.ContactPhone
		}

		address := sql.NullString{}
		if bank.Address != "" {
			address.Valid = true
			address.String = bank.Address
		}
		bankDetails := database.CreateBankParams{
			ID:              uuid.New(),
			CreatedAt:       time.Now().UTC(),
			UpdatedAt:       time.Now().UTC(),
			DeletedAt:       sql.NullTime{},
			IsDeleted:       false,
			Name:            bank.Name,
			ContactPhone:    contactPhone,
			ContactEmail:    bank.ContactEmail,
			Address:         address,
			Type:            bank.Type,
			EstablishedYear: bank.EstablishedYear,
		}
		_, err := DB.GetBankByName(context.Background(), bankDetails.Name)
		if err == nil {
			// bank already created
			continue
		}

		_, err = DB.CreateBank(context.Background(), bankDetails)
		if err != nil {
			log.Fatalln("Not Able to create bank details:", bankDetails.Name, err)
		}
	}
}

func ParseBanksJSON(filename string) ([]CreateBankParams, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Read the file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Parse the JSON into a slice of CreateBankParams
	var banks []CreateBankParams
	if err := json.Unmarshal(bytes, &banks); err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %v", err)
	}

	return banks, nil
}
