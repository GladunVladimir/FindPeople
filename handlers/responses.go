package handlers

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid input"`
}

type MessageResponse struct {
	Message string `json:"message" example:"Friends linked"`
}
