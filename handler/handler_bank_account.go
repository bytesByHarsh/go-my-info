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

func AddAccount(w http.ResponseWriter, r *http.Request, user database.User) {
	params := models.AddBankAccountReq{}

	err := models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	dbBankAccount, err := apiCfg.DB.CreateBankAccount(r.Context(), database.CreateBankAccountParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		DeletedAt:     sql.NullTime{},
		IsDeleted:     false,
		Name:          params.Name,
		BankID:        params.BankID,
		AccountNumber: params.AccountNumber,
		AccountType:   database.BankAccountType(params.AccountType),
		Balance:       params.Balance,
		Currency:      params.Currency,
		UserID:        user.ID,
	})
	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't create bank account: %v", err),
		)
		return
	}
	resp := models.JSONResp{
		Status:  "success",
		Message: "Bank Account Created",
		Data:    models.ConvAccountToAccount(dbBankAccount),
	}
	responseWithJson(w, 201, resp)
}

func GetAccount(w http.ResponseWriter, r *http.Request, user database.User) {
	account_id_str := chi.URLParam(r, "account_id")
	account_id, err := uuid.Parse(account_id_str)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			"Incorrect Account Id provided",
		)
		return
	}

	dbAccount, err := apiCfg.DB.GetBankAccountById(r.Context(), account_id)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("couldn't get account details: %v", err),
		)
		return
	}
	responseWithJson(w, http.StatusOK, models.ConvAccountToAccount(dbAccount))
}

func UpdateAccount(w http.ResponseWriter, r *http.Request, user database.User) {
	account_id_str := chi.URLParam(r, "account_id")
	account_id, err := uuid.Parse(account_id_str)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			"Incorrect Account Id provided",
		)
		return
	}

	params := models.UpdateBankAccountReq{}

	err = models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	apiCfg.DB.UpdateBankAccount(r.Context(), database.UpdateBankAccountParams{
		ID:            account_id,
		UpdatedAt:     time.Now().UTC(),
		Name:          params.Name,
		AccountNumber: params.AccountNumber,
		AccountType:   database.BankAccountType(params.AccountType),
		Balance:       params.Balance,
		Currency:      params.Currency,
	})

	resp := models.JSONResp{
		Status:  "success",
		Message: "Account Data Updated",
		Data:    nil,
	}
	responseWithJson(w, http.StatusAccepted, resp)
}

func GetAllAccounts(w http.ResponseWriter, r *http.Request, user database.User) {
	page, items_per_page, err := parsePaginatedReq(r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("incorrect data: %v", err),
		)
		return
	}

	dbBankAccountList, err := apiCfg.DB.GetAllBankAccount(r.Context(), database.GetAllBankAccountParams{
		Limit:  int32(items_per_page),
		Offset: int32((page - 1) * items_per_page),
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("couldn't get account list: %v", err),
		)
		return
	}

	total_count, err := apiCfg.DB.GetBankAccountCount(r.Context())
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("couldn't get account list: %v", err),
		)
		return
	}
	resp := models.PaginatedListResp[models.BankAccount]{
		Data:         models.CreateAccountListResp(dbBankAccountList),
		Page:         page,
		ItemsPerPage: items_per_page,
		TotalCount:   int(total_count),
	}
	resp.UpdateHasMore()
	responseWithJson(w, http.StatusOK, resp)
}
