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

// AddAccount godoc
//
//	@Summary		Create Account
//	@Description	create new bank account
//	@Tags			Bank Accounts
//	@Accept			json
//	@Produce		json
//	@Param			account	body		models.AddBankAccountReq	true	"Bank Account Body"
//	@Success		201		{object}	models.BankAccount
//	@Failure		400		{object}	models.JSONerrResponse
//	@Router			/accounts/ [post]
func AddAccount(w http.ResponseWriter, r *http.Request, user database.User) {
	params := models.AddBankAccountReq{}

	err := models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	bankId, err := uuid.Parse(params.BankID)

	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing Bank UUID: %v", err),
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
		BankID:        bankId,
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

// GetAccount godoc
//
//	@Summary		Get Account
//	@Description	get bank account details
//	@Tags			Bank Accounts
//	@Accept			json
//	@Produce		json
//	@Param			account_id	path		string	true	"Bank Account ID"
//	@Success		201			{object}	models.BankAccount
//	@Failure		400			{object}	models.JSONerrResponse
//	@Router			/accounts/{account_id} [get]
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

// UpdateAccount godoc
//
//	@Summary		Update Account
//	@Description	update bank account details
//	@Tags			Bank Accounts
//	@Accept			json
//	@Produce		json
//	@Param			account_id	path		string						true	"Bank Account ID"
//	@Param			account		body		models.UpdateBankAccountReq	true	"Bank Account Body"
//	@Success		201			{object}	models.JSONerrResponse
//	@Failure		400			{object}	models.JSONerrResponse
//	@Router			/accounts/{account_id} [put]
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

// GetAllAccounts godoc
//
//	@Summary		Get All Account
//	@Description	get all bank account
//	@Tags			Bank Accounts
//	@Accept			json
//	@Produce		json
//	@Param			page			query		int32	true	"Page Number"
//	@Param			items_per_page	query		int32	true	"Items Per Page"
//	@Success		200				{object}	models.PaginatedListResp[models.BankAccount]
//	@Failure		400				{object}	models.JSONerrResponse
//	@Router			/accounts/ [get]
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
