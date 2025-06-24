package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "supersecret"

func hashPassword(password string) (string, error) {

	// use bcrypt library to hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err

}

func CheckPassword(password string, hashedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateJWTToken(email string) (string, error) {

	// with jwt claims simply mean the data that's attached to it.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(2 * time.Hour).Unix(),
	})

	// key should be byte slice
	return token.SignedString([]byte(secretKey))
}
