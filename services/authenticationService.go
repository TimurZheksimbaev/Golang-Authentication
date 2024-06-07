package services

import (
	"github.com/TimurZheksimbaev/Golang-auth/models"
	"github.com/TimurZheksimbaev/Golang-auth/network/requests"
	"github.com/TimurZheksimbaev/Golang-auth/storage"
	"github.com/TimurZheksimbaev/Golang-auth/utils"
	"github.com/go-playground/validator/v10"
)

type AuthService struct {
	UserStorage storage.UserStorageI
	Validate *validator.Validate
}

func NewAuthService(us storage.UserStorageI, validate *validator.Validate) AuthServiceI {
	return &AuthService{
		UserStorage: us,
		Validate: validate,
	}
}

func (as *AuthService) Login(loginRequest requests.LoginRequest) (string, error) {
	// find user in database
	user, err := as.UserStorage.FindByEmail(loginRequest.Email)
	if err != nil {
		return "", utils.ServiceError("Error finding such user", err)
	}
	
	// compare passwords
	err = utils.ComparePasswords(user.Password, loginRequest.Password)
	if err != nil {
		return "", utils.ServiceError("Error finding such user", err)
	}

	// generate token
	tokenPayload := utils.TokenPayload{
		Email: user.Email,
		Password: user.Password,
	}
	token, err := utils.GenerateToken(tokenPayload)
	if err != nil {
		return "", utils.ServiceError("Could not receive token", err)
	}

	return token, err
}

func (as *AuthService) Register(registerRequest requests.RegisterRequest) error {

	// hash password
	hashedPassword, err := utils.EncryptPassword(registerRequest.Password)
	if err != nil {
		return utils.ServiceError("Could not receive hashed password", err)
	}

	// construct user model
	newUser := models.User{
		Username: registerRequest.Username,
		Email: registerRequest.Email,
		Password: hashedPassword,
	}

	// create a new entry in database
	err = as.UserStorage.Create(&newUser)
	return err
}