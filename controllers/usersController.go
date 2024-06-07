package controllers

import (
	"net/http"
	"github.com/TimurZheksimbaev/Golang-auth/network/responses"
	"github.com/TimurZheksimbaev/Golang-auth/storage"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	userStorage storage.UserStorageI
}

func NewUsersController(us storage.UserStorageI) *UsersController {
	return &UsersController{
		userStorage: us,
	}
}

func (uc *UsersController) GetUsers(ctx *gin.Context) {
	users, err := uc.userStorage.GetAll()
	if err != nil {
		response := responses.Response{
			Code: http.StatusInternalServerError,
			Status: "Internal error",
			Message: "Error fetching all users",
		}
		ctx.JSON(http.StatusInternalServerError, response)
	}

	webResponse := responses.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)

}