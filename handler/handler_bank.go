package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/bytesByHarsh/go-my-info/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func CreateBank(w http.ResponseWriter, r *http.Request, user database.User) {
	if user.Role != UserRole_Admin {
		responseWithError(w, http.StatusUnauthorized,
			"Access Denied",
		)
		return
	}
	params := models.CreateBankReq{}

	err := models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}
	contactPhone := sql.NullString{}
	if params.ContactPhone != "" {
		contactPhone.Valid = true
		contactPhone.String = params.ContactPhone
	}

	address := sql.NullString{}
	if params.Address != "" {
		address.Valid = true
		address.String = params.Address
	}

	dbBank, err := apiCfg.DB.CreateBank(r.Context(), database.CreateBankParams{
		ID:              uuid.New(),
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
		DeletedAt:       sql.NullTime{},
		IsDeleted:       false,
		Name:            params.Name,
		ContactPhone:    contactPhone,
		ContactEmail:    params.ContactEmail,
		Address:         address,
		Type:            params.BankType,
		EstablishedYear: params.EstablishedYear,
	})

	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't create bank: %v", err),
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "User Created",
		Data:    models.ConvBankToBank(dbBank),
	}
	responseWithJson(w, 201, resp)
}

func GetBankList(w http.ResponseWriter, r *http.Request) {
	page, items_per_page, err := parsePaginatedReq(r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("incorrect data: %v", err),
		)
		return
	}

	dbBankList, err := apiCfg.DB.GetAllBank(r.Context(), database.GetAllBankParams{
		Limit:  int32(items_per_page),
		Offset: int32((page - 1) * items_per_page),
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("couldn't get user list: %v", err),
		)
		return
	}
	total_count, err := apiCfg.DB.GetBankCount(r.Context())
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("couldn't get user list: %v", err),
		)
		return
	}
	resp := models.PaginatedListResp[models.Bank]{
		Data:         models.CreateBankListResp(dbBankList),
		Page:         page,
		ItemsPerPage: items_per_page,
		TotalCount:   int(total_count),
	}
	resp.UpdateHasMore()
	responseWithJson(w, http.StatusOK, resp)
}

func UpdateBank(w http.ResponseWriter, r *http.Request, user database.User) {
	if user.Role != UserRole_Admin {
		responseWithError(w, http.StatusUnauthorized,
			"Access Denied",
		)
		return
	}
	bankId := chi.URLParam(r, "bank_id")
	id, err := uuid.Parse(bankId)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			"Incorrect Bank Id provided",
		)
		return
	}
	params := models.UpdateBankReq{}

	err = models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	contactPhone := sql.NullString{}
	if params.ContactPhone != "" {
		contactPhone.Valid = true
		contactPhone.String = params.ContactPhone
	}

	address := sql.NullString{}
	if params.Address != "" {
		address.Valid = true
		address.String = params.Address
	}

	err = apiCfg.DB.UpdateBank(r.Context(), database.UpdateBankParams{
		ID:           id,
		UpdatedAt:    time.Now().UTC(),
		Name:         params.Name,
		ContactPhone: contactPhone,
		ContactEmail: params.ContactEmail,
		Address:      address,
		Type:         params.BankType,
	})

	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't update bank: %v", err),
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "Bank Data Updated",
		Data:    nil,
	}
	responseWithJson(w, http.StatusAccepted, resp)
}
