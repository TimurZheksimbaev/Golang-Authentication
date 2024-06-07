package utils

import (
	"fmt"
	"time"

	"github.com/TimurZheksimbaev/Golang-auth/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	Email string
	Password string
}

func GenerateToken(payload TokenPayload) (string, error) {
	appConfig, err := config.LoadEnv()
	
	if err != nil {
		return "", ConfigError("Could not read config", err)
	}

	ttl := appConfig.TokenExpiresIn
	secretKey := appConfig.TokenSecret

	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now()
	claim := token.Claims.(jwt.MapClaims)
	claim["sub"] = payload
	claim["exp"] = now.Add(ttl).Unix()
	claim["iat"] = now.Unix()
	claim["nbf"] = now.Unix()
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", TokenError("Could not generate token", err)
	}
	return tokenString, nil
}

func ValidateToken(token string, signedKey string) (any, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected method %s", jwtToken.Header["alg"])
		}
		return []byte(signedKey), nil
	})
	if err != nil {
		return nil, TokenError("Could not validate parse token", err)
	}
	claim, ok := tok.Claims.(jwt.MapClaims)
	if !ok  || !tok.Valid {
		return nil, TokenError("Invalid token claims", fmt.Errorf("wrong claim %s", ok))
	}
	return claim["sub"], nil
}