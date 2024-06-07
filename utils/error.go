package utils

import (
	"fmt"
	"log"
)

// All errors

func Log(err error) {
	if err != nil {
		log.Println(err)
	}
}

func LogExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}



func DatabaseError(message string, err error) error {
	return fmt.Errorf("[DATABASE] %s: %s", message, err)
}

func ConfigError(message string, err error) error {
	return fmt.Errorf("[CONFIG] %s: %s", message, err)
}

func ServiceError(message string, err error) error {
	return fmt.Errorf("[SERVICE] %s: %s", message, err)
}

func ControllerError(message string, err error) error {
	return fmt.Errorf("[CONTROLLER] %s: %s", message, err)
}

func TokenError(message string, err error) error {
	return fmt.Errorf("[TOKEN] %s: %s", message, err)
}

func PasswordError(message string, err error) error {
	return fmt.Errorf("[PASSWORD] %s: %s", message, err)
}