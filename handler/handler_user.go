package handler

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/bytesByHarsh/go-my-info/config"
	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/bytesByHarsh/go-my-info/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const (
	UserRole_Admin = 100
	UserRole_Base  = 10
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	params := models.CreateUserReq{}

	err := models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	hash := HashPassword(params.Password)

	dbUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		DeletedAt:      sql.NullTime{},
		IsDeleted:      false,
		Name:           params.Name,
		Email:          params.Email,
		Username:       params.Username,
		PhoneNum:       "",
		ProfileImg:     "",
		Role:           10,
		HashedPassword: hash,
	})

	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't create user: %v", err),
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "User Created",
		Data:    models.ConvUserToUser(dbUser),
	}
	responseWithJson(w, 201, resp)
}

func CreateUserByAdmin(w http.ResponseWriter, r *http.Request, user database.User) {
	if user.Role != UserRole_Admin {
		responseWithError(w, http.StatusUnauthorized,
			"Access Denied",
		)
		return
	}

	params := models.CreateUserByAdminReq{}

	err := models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	hash := HashPassword(params.Password)

	var role int32
	if params.IsSuperUser {
		role = 100
	} else {
		role = 10
	}

	dbUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		DeletedAt:      sql.NullTime{},
		IsDeleted:      false,
		Name:           params.Name,
		Email:          params.Email,
		Username:       params.Username,
		PhoneNum:       "",
		ProfileImg:     "",
		Role:           role,
		HashedPassword: hash,
	})

	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't create user: %v", err),
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "User Created",
		Data:    models.ConvUserToUser(dbUser),
	}
	responseWithJson(w, 201, resp)

}

func GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJson(w, http.StatusOK, models.ConvUserToUser(user))
}

func GetAnotherUser(w http.ResponseWriter, r *http.Request, user database.User) {
	username := chi.URLParam(r, "username")
	userDb, err := apiCfg.DB.GetUserByUsername(r.Context(), username)
	if err != nil {
		responseWithError(w, http.StatusNotFound,
			fmt.Sprintf("Couldn't get user: %v", err),
		)
		return
	}
	responseWithJson(w, http.StatusOK, models.ConvUserToUser(userDb))
}

func GetUserList(w http.ResponseWriter, r *http.Request, user database.User) {
	if user.Role != UserRole_Admin {
		responseWithError(w, http.StatusUnauthorized,
			"Access Denied",
		)
		return
	}

	page, items_per_page, err := parsePaginatedReq(r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("incorrect data: %v", err),
		)
		return
	}

	dbUserList, err := apiCfg.DB.GetAllUsers(r.Context(), database.GetAllUsersParams{
		Limit:  int32(items_per_page),
		Offset: int32((page - 1) * items_per_page),
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("couldn't get user list: %v", err),
		)
		return
	}

	total_count, err := apiCfg.DB.GetUserCount(r.Context())
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("couldn't get user list: %v", err),
		)
		return
	}

	resp := models.PaginatedListResp[models.User]{
		Data:         models.CreateUserListResp(dbUserList),
		Page:         page,
		ItemsPerPage: items_per_page,
		TotalCount:   int(total_count),
	}
	resp.UpdateHasMore()
	responseWithJson(w, http.StatusOK, resp)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, user database.User) {
	params := models.UpdateUserReq{}

	err := models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	err = apiCfg.DB.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:         user.ID,
		UpdatedAt:  time.Now().UTC(),
		Name:       params.Name,
		PhoneNum:   params.PhoneNum,
		Email:      params.Email,
		Username:   params.Username,
		ProfileImg: params.ProfileImg,
		Role:       user.Role,
	})

	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't update user: %v", err),
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "User Data Updated",
		Data:    nil,
	}
	responseWithJson(w, http.StatusAccepted, resp)
}
func UpdateAnotherUser(w http.ResponseWriter, r *http.Request, user database.User) {
	if user.Role != UserRole_Admin {
		responseWithError(w, http.StatusUnauthorized,
			"Access Denied",
		)
		return
	}
	userId := chi.URLParam(r, "user_id")
	id, err := uuid.Parse(userId)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			"Incorrect User Id provided",
		)
		return
	}

	params := models.UpdateUserReq{}

	err = models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	err = apiCfg.DB.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:         id,
		UpdatedAt:  time.Now().UTC(),
		Name:       params.Name,
		PhoneNum:   params.PhoneNum,
		Email:      params.Email,
		Username:   params.Username,
		ProfileImg: params.ProfileImg,
		Role:       user.Role,
	})

	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't update user: %v", err),
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "User Data Updated",
		Data:    nil,
	}
	responseWithJson(w, http.StatusAccepted, resp)
}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request, user database.User) {
	params := models.UpdatePasswordReq{}

	err := models.VerifyJson(&params, r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	err = apiCfg.DB.UpdateUserPassword(r.Context(), database.UpdateUserPasswordParams{
		ID:             user.ID,
		UpdatedAt:      time.Now().UTC(),
		HashedPassword: HashPassword(params.Password),
	})

	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't update user password: %v", err),
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "User Password Updated",
		Data:    nil,
	}
	responseWithJson(w, http.StatusAccepted, resp)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, user database.User) {
	err := apiCfg.DB.DeleteUser(r.Context(), database.DeleteUserParams{
		ID:        user.ID,
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't delete user: %v", err),
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "User Deleted",
		Data:    nil,
	}
	responseWithJson(w, http.StatusAccepted, resp)
}

func DbDeleteUser(w http.ResponseWriter, r *http.Request, user database.User) {
	if user.Role < UserRole_Admin {
		responseWithError(w, http.StatusUnauthorized,
			"Proper Authentication Required",
		)
		return
	}
	err := apiCfg.DB.HardDeleteUser(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't permanently delete user: %v", err),
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "User Deleted Permanently",
		Data:    nil,
	}
	responseWithJson(w, http.StatusAccepted, resp)
}

func HashPassword(password string) string {
	// bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// return string(bytes), err
	// Concatenate the secret and password
	combined := config.Cfg.SECRET_KEY + password
	hash := sha256.New()
	hash.Write([]byte(combined))
	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}
