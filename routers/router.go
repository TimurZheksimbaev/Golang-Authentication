package routers

import (
	"net/http"

	"github.com/TimurZheksimbaev/Golang-auth/controllers"
	"github.com/TimurZheksimbaev/Golang-auth/storage"
	"github.com/gin-gonic/gin"
)


func NewRouter(authController controllers.AuthController, us storage.UserStorageI, usersController controllers.UsersController) *gin.Engine {
	server := gin.Default()

	server.GET("", func (ctx *gin.Context)  {
		ctx.String(http.StatusOK, "Welcome home")
	})

	server.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Page not found")
	})

	router := server.Group("/api")
	authRouter := router.Group("/auth")

	authRouter.POST("/register", authController.Register)
	authRouter.POST("/login", authController.Login)

	usersRouter := router.Group("/users")
	usersRouter.GET("", authController.AuthMiddleware(), usersController.GetUsers)
	return server	
}