package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/bytesByHarsh/go-my-info/models"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	params := models.AuthReq{}
	err := models.VerifyJson(&params, r)

	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	user, err := apiCfg.DB.GetUserByUsername(r.Context(), params.Username)
	if err != nil {
		responseWithError(w, http.StatusNotFound,
			fmt.Sprintf("Couldn't get user: %v", err),
		)
		return
	}

	if !checkPassword(user.HashedPassword, params.Password) {
		responseWithError(w, http.StatusBadRequest,
			"wrong password",
		)
		return
	}

	expirationTime := time.Now().UTC().Add(8 * time.Hour) //Valid for 8hrs
	// expirationTime := time.Now().UTC().Add(time.Second)
	var expirationTimeStr string = strconv.FormatUint(uint64(expirationTime.Unix()), 10)
	claims := map[string]interface{}{
		"username": user.Username,
		"role":     user.Role,
		"time":     expirationTimeStr,
	}

	_, tokenString, err := apiCfg.AuthToken.Encode(claims)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,
			"token not generated",
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "User Logged In",
		Data:    models.AuthResp{Token: tokenString},
	}
	cookie := http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		MaxAge:   8 * 60 * 60,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		// Path:     "/",
		// HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	responseWithJson(w, 201, resp)

}

func MiddlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from the cookie
		c, err := r.Cookie("auth_token")
		if err != nil {
			if err == http.ErrNoCookie {
				responseWithError(w, http.StatusUnauthorized, "User Not Authenticated")
				return
			}
			responseWithError(w, http.StatusBadRequest, "Bad request")
			return
		}

		tokenStr := c.Value

		token, err := apiCfg.AuthToken.Decode(tokenStr)
		if err != nil {
			responseWithError(w, http.StatusBadRequest, "incorrect token")
			return
		}
		claims, err := token.AsMap(r.Context())
		if err != nil {
			responseWithError(w, http.StatusBadRequest, "incorrect token")
			return
		}

		username, _ := claims["username"].(string)
		timeStr, _ := claims["time"].(string)

		expTime, _ := strconv.Atoi(timeStr)
		if uint64(time.Now().UTC().Unix()) > uint64(expTime) {
			responseWithError(w, http.StatusUnauthorized, "token timeout")
			return
		}

		user, err := apiCfg.DB.GetUserByUsername(r.Context(), username)
		if err != nil {
			responseWithError(w, http.StatusNotFound,
				fmt.Sprintf("Couldn't get user: %v", err),
			)
			return
		}

		handler(w, r, user)
	}
}

func checkPassword(hashPass, inputPass string) bool {
	return hashPassword(inputPass) == hashPass
}
