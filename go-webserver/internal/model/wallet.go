package model

type ErrorResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetWalletBalanceResponse struct {
	UserID  string  `json:"user_id"`
	Balance float32 `json:"balance"`
}

type CreateWalletResponse struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

type CreateWalletRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type UpdateWalletResponse struct {
	UserID  string  `json:"user_id"`
	Balance float32 `json:"balance"`
}

type UpdateWalletRequest struct {
	UserID  string  `json:"user_id"`
	Balance float32 `json:"balance" binding:"required"`
}
