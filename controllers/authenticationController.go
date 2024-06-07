package controllers

import (
	"net/http"
	"time"
	"strings"
	"github.com/TimurZheksimbaev/Golang-auth/config"
	"github.com/TimurZheksimbaev/Golang-auth/network/requests"
	"github.com/TimurZheksimbaev/Golang-auth/network/responses"
	"github.com/TimurZheksimbaev/Golang-auth/services"
	"github.com/TimurZheksimbaev/Golang-auth/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthServiceI
	appConfig config.AppConfig
}

func NewAuthController(authService services.AuthServiceI, appConfig config.AppConfig) *AuthController {
	return &AuthController{
		authService: authService,
		appConfig: appConfig,
	}
}

func (ac *AuthController) Login(ctx *gin.Context) {
	// parse login request
	loginRequest := requests.LoginRequest{}
	err := ctx.ShouldBindBodyWithJSON(&loginRequest)
	utils.Log(err)


	// get token
	token, tokenErr := ac.authService.Login(loginRequest)
	if tokenErr != nil {
		response := responses.Response{
			Code: http.StatusBadRequest,
			Status: "Bad request",
			Message: "Invalid password or email",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return 
	}

	// construct login response with token
	loginResponse := responses.LoginResponse{
		TokenType: "Bearer",
		Token: token,
		TokenCreatedAt: time.Now(),
	}

	// construct response
	response := responses.Response{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Successfully logged in",
		Data: loginResponse,
	}
	ctx.SetCookie("token", token, ac.appConfig.TokenMaxAge*60, "/", "localhost", false, true)
	// send response
	ctx.JSON(http.StatusOK, response)
}

func (ac *AuthController) Register(ctx *gin.Context) {
	// parse register request
	registerRequest := requests.RegisterRequest{}
	err := ctx.ShouldBindBodyWithJSON(&registerRequest)
	utils.Log(err)

	// create a new user
	err = ac.authService.Register(registerRequest)
	if err != nil {
		response := responses.Response{
			Code: http.StatusBadRequest,
			Status: "Bad request",
			Message: "Invalid password or login or username",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// construct and send response
	response := responses.Response{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Successfully registered",
	}

	ctx.JSON(http.StatusOK, response)
}

func (ac *AuthController) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the token from the Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "Authorization header is missing",
			})
			return
		}

		// Check if the header is in the expected format (Bearer token)
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "Invalid authorization header format",
			})
			return
		}

		// Validate the token
		tokenString := parts[1]
		claims, err := utils.ValidateToken(tokenString, ac.appConfig.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "Invalid or expired token",
			})
			return
		}

		// Set the user ID in the context for further use in handlers
		ctx.Set("email", claims)

		// Continue to the next handler
		ctx.Next()
	}
}