package main

import (
	"authenticate-service/data"
	"authenticate-service/logs"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Request models
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"test@test.com"`
	Password string `json:"password" binding:"required" example:"tester"`
}
type SignupRequest struct {
	Email    string `json:"email" binding:"required,email" example:"test@test.com"`
	Username string `json:"username" binding:"required" example:"testuser"`
	Password string `json:"password" binding:"required" example:"tester"`
}

// Response models
type LoginResponse struct {
	Error   bool   `json:"error" example:"false"`
	Message string `json:"message" example:"Authentication successful"`
	Email   string `json:"email" example:""`
	Token   string `json:"token" example:""`
}
type SignupResponse struct {
	Error   bool   `json:"error" example:"false"`
	Message string `json:"message" example:"User created successfully"`
}

type ErrorResponse struct {
	Error   bool   `json:"error" example:"true"`
	Message string `json:"message" example:"Error message"`
}

// @Summary Authenticate user
// @Description Authenticate user with email and password
// @Tags Authentication
// @version 1.0
// @produce application/json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse "Bad Request - Invalid request body"
// @Failure 401 {object} ErrorResponse "Unauthorized - Invalid credentials"
// @Failure 500 {object} ErrorResponse "Internal Server Error - JWT generation failed"
// @Router /v1/authenticate/login [post]
func (app *Config) Authenticate(c *gin.Context) {

	var request LoginRequest
	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   true,
			Message: "Failed to parse request body",
		})
		return
	}

	// start validate credentials
	user, err := app.Repo.GetByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   true,
			Message: "Invalid email or password",
		})
		return
	}

	isAuthenticated := CheckPassword(request.Password, user.Password)

	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   true,
			Message: "Invalid email or password",
		})
		return
	}

	// generate JWT token
	token, err := generateJWTToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   true,
			Message: "Failed to generate JWT token",
		})
		return
	}

	// call logger service to log the authentication event
	// err = app.logAuthenticationEvent("authenticate-service", "info", fmt.Sprintf("User %s logged in at %v", user.Email, time.Now().Format(time.RFC3339)))
	err = app.logAuthenticationEventViaGRPC("authenticate-service", "info", fmt.Sprintf("User %s logged in at %v", user.Email, time.Now().Format(time.RFC3339)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   true,
			Message: "Failed to log authentication event",
		})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Error:   false,
		Message: "Authentication successful",
		Email:   user.Email,
		Token:   token,
	})
}

// @Summary Register a new user
// @Description Register user with email, username, and password
// @Tags Authentication
// @version 1.0
// @produce application/json
// @Param request body SignupRequest true "Registration credentials"
// @Success 200 {object} SignupResponse
// @Failure 400 {object} ErrorResponse "Bad Request - Invalid request body"
// @Failure 401 {object} ErrorResponse "Unauthorized - Failed to hash the password"
// @Failure 500 {object} ErrorResponse "Internal Server Error - Failed to register user"
// @Router /v1/authenticate/signup [post]
func (app *Config) Signup(c *gin.Context) {

	var request SignupRequest
	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   true,
			Message: "Failed to parse request body",
		})
		return
	}

	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   true,
			Message: "Failed to hash the password",
		})
		return
	}

	_, err = app.Repo.Insert(data.User{
		Email:    request.Email,
		Username: request.Username,
		Password: hashedPassword,
		Role:     "guest", // default role,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   true,
			Message: "Failed to register new user",
		})
		return
	}

	c.JSON(http.StatusCreated, SignupResponse{
		Error:   false,
		Message: "User created successfully",
	})
}

func (app *Config) logAuthenticationEvent(name, level, message string) error {

	var logPayload struct {
		Name    string `json:"name"`
		Level   string `json:"level"`
		Message string `json:"message"`
	}

	logPayload.Name = name
	logPayload.Level = level
	logPayload.Message = message

	// Prepare the request to the logger service
	jsonData, _ := json.MarshalIndent(logPayload, "", "\t")

	// Here will call the remote logger service which isn't running when testing.
	// How to mock the request?
	logServiceURL := "http://logger-service/v1/log"

	// register the http request
	req, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// Register the http client and send the request
	// client := &http.Client{}
	_, err = app.Clients.LogHTTPClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (app *Config) logAuthenticationEventViaGRPC(name, level, message string) error {

	/*
		Register a grpc client with grpc.NewClient method. When you writing the handler unit test, it is hard to mock the grpc client.
		Therefore we need to use dependency injection to inject the grpc client into the handler.
		In order to do this, we need to define an interface for the grpc client and implement it in a struct.
		Then we can mock the grpc client in the unit test.

		conn, err := grpc.NewClient("logger-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}
		defer conn.Close()
		c := logs.NewLogServiceClient(conn)
	*/

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := app.Clients.LoggRPCClient.WriteLog(ctx, &logs.LogRequest{
		LogEntry: &logs.Log{
			Name:    name,
			Level:   level,
			Message: message,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
