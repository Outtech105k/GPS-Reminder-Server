package responses

type SuccessDefaultResponse struct {
	Message string `json:"message"`
}

type ErrorDefaultResponse struct {
	Error string `json:"error"`
}
