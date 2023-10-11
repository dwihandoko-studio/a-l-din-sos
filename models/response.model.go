package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Token struct {
	Access     *string `json:"access_token"`
	Expired_in int64   `json:"expired_in"`
	Refresh    *string `json:"refresh_token"`
}
