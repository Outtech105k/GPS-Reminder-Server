package responses

type SuccessTokenResponse struct {
	Code   int    `json:"code"`
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

type ErrorTokenResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
