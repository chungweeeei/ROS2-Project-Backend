package main

import (
	"auth-service/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"test@test.com"`
	Password string `json:"password" binding:"required" example:"12345678"`
}

type SignupRequest struct {
	Email     string `json:"email" binding:"required,email" example:"test@test.com"`
	FirstName string `json:"first_name" binding:"required" example:"Andy"`
	LastName  string `json:"last_name" binding:"required" example:"Tseng"`
	Password  string `json:"password" binding:"required" example:"12345678"`
}

func (app *Config) Authenticate(c *gin.Context) {

	var request LoginRequest
	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse request body",
		})
		return
	}

	// validate credentials
	user, err := app.Models.User.GetByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	isAuthenticated := app.Models.User.PasswordMatches(request.Password, user.Password)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	token, err := generateJWTToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate JWT token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Authenticate successfully",
		"email":   user.Email,
		"token":   token,
	})
}

func (app *Config) Signup(c *gin.Context) {

	var request SignupRequest
	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		app.ErrorLog.Println("Error binding request:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse request body",
		})
		return
	}

	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash the password",
		})
		return
	}

	_, err = app.Models.User.Insert(data.User{
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  hashedPassword,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register new user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}
