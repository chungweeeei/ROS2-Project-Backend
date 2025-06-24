package main

import (
	"authenticate-service/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Config) Authenticate(c *gin.Context) {

	var loginRequest struct {
		Email    string `json:"email";binding:"required,email"`
		Password string `json:"password";binding:"required"`
	}

	err := c.ShouldBindBodyWithJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to parse request body",
		})
		return
	}

	// start validate credentials
	user, err := app.Models.User.GetByEmail(loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Invalid email or password",
		})
		return
	}

	isAuthenticated := CheckPassword(loginRequest.Password, user.Password)

	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Invalid email or password",
		})
		return
	}

	// generate JWT token
	token, err := generateJWTToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to generate JWT token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Authentication successful",
		"email":   user.Email,
		"token":   token,
	})
}

func (app *Config) Signup(c *gin.Context) {

	var userRequest struct {
		Email    string `json:"email";binding:"required,email"`
		Username string `json:"username";binding:"required"`
		Password string `json:"password";binding:"required"`
	}

	err := c.ShouldBindBodyWithJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to parse request body",
		})
		return
	}

	hashedPassword, err := hashPassword(userRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to hash the password",
		})
		return
	}

	err = app.Models.User.Insert(data.User{
		Email:    userRequest.Email,
		Username: userRequest.Username,
		Password: hashedPassword,
		Role:     "guest", // default role,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to register new user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}
