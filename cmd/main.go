package main

import (

	"github.com/TimurZheksimbaev/Golang-auth/config"
	"github.com/TimurZheksimbaev/Golang-auth/controllers"
	"github.com/TimurZheksimbaev/Golang-auth/models"
	"github.com/TimurZheksimbaev/Golang-auth/routers"
	"github.com/TimurZheksimbaev/Golang-auth/services"
	"github.com/TimurZheksimbaev/Golang-auth/storage"
	"github.com/TimurZheksimbaev/Golang-auth/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	// config
	appConfig, err := config.LoadEnv()
	utils.LogExit(err)

	// database init
	db, err := config.ConnectToDB(appConfig)
	utils.LogExit(err)
	db.Table("users").AutoMigrate(&models.User{})

	// construct storage
	storage := storage.NewStorage(db)

	// service and controller
	authService := services.NewAuthService(storage, validator.New())
	authController := controllers.NewAuthController(authService, *appConfig)
	usersController := controllers.NewUsersController(storage)

	router := routers.NewRouter(*authController, storage, *usersController)
	server := &http.Server{
		Addr:           appConfig.ServerHost + ":" + fmt.Sprint(appConfig.ServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	utils.LogExit(server.ListenAndServe())
}

