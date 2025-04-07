package handler

import (
	"net/http"

	"example.com/webserver/internal/model"
	"example.com/webserver/internal/wallet"
	"github.com/gin-gonic/gin"
)

func GetWalletBalanceHandler(c *gin.Context) {
	userID := c.Param("userId")
	balance, err := wallet.GetWalletBalance(userID)
	if err != nil {
		errResponse := model.ErrorResponse{
			Status:  "error",
			Code:    "wallet_not_found",
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, errResponse)
		return
	}
	response := model.GetWalletBalanceResponse{
		UserID:  userID,
		Balance: balance,
	}
	c.JSON(http.StatusOK, response)
}

func CreateWalletHandler(c *gin.Context) {
	request := model.CreateWalletRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errResponse := model.ErrorResponse{
			Status:  "error",
			Code:    "invalid_request",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	err = wallet.CreateWallet(request.UserID)
	if err != nil {
		errResponse := model.ErrorResponse{
			Status:  "error",
			Code:    "wallet_already_exists",
			Message: err.Error(),
		}
		c.JSON(http.StatusConflict, errResponse)
		return
	}

	response := model.CreateWalletResponse{
		UserID:  request.UserID,
		Message: "Wallet created successfully",
	}
	c.JSON(http.StatusCreated, response)
}

func UpdateWalletHandler(c *gin.Context) {
	request := model.UpdateWalletRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errResponse := model.ErrorResponse{
			Status:  "error",
			Code:    "invalid_request",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	balance, err := wallet.UpdateWalletBalance(request.UserID, request.Balance)
	if err != nil {
		errResponse := model.ErrorResponse{
			Status:  "error",
			Code:    "wallet_not_found",
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, errResponse)
		return
	}
	response := model.UpdateWalletResponse{
		UserID:  request.UserID,
		Balance: balance,
	}
	c.JSON(http.StatusOK, response)
}
