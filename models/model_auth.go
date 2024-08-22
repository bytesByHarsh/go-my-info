package models

type AuthResp struct {
	Token string `json:"token"`
}

type AuthReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
