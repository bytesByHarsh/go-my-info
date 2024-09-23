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

func AddCard(w http.ResponseWriter, r *http.Request, user database.User) {
	params := models.AddCardReq{}

	err := models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	var account_id uuid.NullUUID
	if params.BankAccountID == "" {
		account_id.Valid = false

	} else {
		account_id.Valid = true
		account_id.UUID, err = uuid.Parse(params.BankAccountID)
		if err != nil {
			responseWithError(w, http.StatusBadRequest,
				"Incorrect Account Id provided",
			)
			return
		}
	}

	dbCard, err := apiCfg.DB.CreateCard(r.Context(), database.CreateCardParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		DeletedAt:     sql.NullTime{},
		IsDeleted:     false,
		BankID:        params.BankID,
		UserID:        user.ID,
		BankAccountID: account_id,
		Name:          params.Name,
		Nickname:      params.Nickname,
		Number:        params.Number,
		Type:          database.CardType(params.CardType),
		ExpMonth:      params.ExpirationMonth,
		ExpYear:       params.ExpirationYear,
		Cvv:           params.Cvv,
		TotalLimit:    params.TotalLimit,
		BillDate:      params.BillDate,
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
		Data:    models.ConvCardToCard(dbCard),
	}
	responseWithJson(w, 201, resp)
}

func GetCard(w http.ResponseWriter, r *http.Request, user database.User) {
	card_id_str := chi.URLParam(r, "card_id")
	card_id, err := uuid.Parse(card_id_str)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			"Incorrect Card Id provided",
		)
		return
	}

	dbCard, err := apiCfg.DB.GetCardById(r.Context(), card_id)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("couldn't get card details: %v", err),
		)
		return
	}
	responseWithJson(w, http.StatusOK, models.ConvCardToCard(dbCard))
}

func UpdateCard(w http.ResponseWriter, r *http.Request, user database.User) {
	card_id_str := chi.URLParam(r, "card_id")
	card_id, err := uuid.Parse(card_id_str)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			"Incorrect Card Id provided",
		)
		return
	}

	params := models.UpdateCardReq{}

	err = models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	var account_id uuid.NullUUID
	if params.BankAccountID == "" {
		account_id.Valid = false

	} else {
		account_id.Valid = true
		account_id.UUID, err = uuid.Parse(params.BankAccountID)
		if err != nil {
			responseWithError(w, http.StatusBadRequest,
				"Incorrect Account Id provided",
			)
			return
		}
	}

	apiCfg.DB.UpdateCard(r.Context(), database.UpdateCardParams{
		ID:        card_id,
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Nickname:  params.Nickname,
		IsActive:  true,
		Number:    params.Number,
		Type:      database.CardType(params.CardType),
		ExpMonth:  params.ExpirationMonth,
		ExpYear:   params.ExpirationYear,
		Cvv:       params.Cvv,
	})
	resp := models.JSONResp{
		Status:  "success",
		Message: "Card Data Updated",
		Data:    nil,
	}
	responseWithJson(w, http.StatusAccepted, resp)
}

func GetAllCards(w http.ResponseWriter, r *http.Request, user database.User) {
	page, items_per_page, err := parsePaginatedReq(r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("incorrect data: %v", err),
		)
		return
	}

	dbCardList, err := apiCfg.DB.GetUserCards(r.Context(), database.GetUserCardsParams{
		Limit:  int32(items_per_page),
		Offset: int32((page - 1) * items_per_page),
		UserID: user.ID,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("couldn't get card list: %v", err),
		)
		return
	}

	total_count, err := apiCfg.DB.GetUserCardCount(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("couldn't get account list: %v", err),
		)
		return
	}
	resp := models.PaginatedListResp[models.Card]{
		Data:         models.CreateCardListResp(dbCardList),
		Page:         page,
		ItemsPerPage: items_per_page,
		TotalCount:   int(total_count),
	}
	resp.UpdateHasMore()
	responseWithJson(w, http.StatusOK, resp)
}
