package main

import (
	"errors"
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

func generateJWTToken(email string) (string, error) {

	// with jwt claims simply mean the data that's attached to it.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(2 * time.Hour).Unix(),
	})

	// key should be byte slice
	return token.SignedString([]byte(secretKey))
}

func verifyToken(token string) (string, error) {

	// keyFn determine how to verify the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// check signing method is HMAC
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		// return verified signing key
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return "", errors.New("token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("could not parse claims")
	}

	return claims["email"].(string), nil
}
