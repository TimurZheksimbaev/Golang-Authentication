package responses

import "time"

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token string `json:"token"`
	TokenCreatedAt time.Time `json:"token_created_at"`
}

type Response struct {
	Code int	`json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}