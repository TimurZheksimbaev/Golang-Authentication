package services

import (
	"github.com/TimurZheksimbaev/Golang-auth/network/requests"
)

type AuthServiceI interface {
	Register(requests.RegisterRequest) error
	Login(requests.LoginRequest) (string, error)
}
