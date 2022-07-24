package login

type GetTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetIHduTokenResponse struct {
	TpUp string `json:"tpUp"`
}
type GetNewJWTokenResponse struct {
	JSESSIONID string `json:"JSESSIONID"`
	Route      string `json:"route"`
}

type GetSKLokenResponse struct {
	Token string `json:" token"`
}
