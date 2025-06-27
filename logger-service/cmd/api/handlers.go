package main

import (
	"logger-service/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request models
type WriteLogRequest struct {
	Name    string `json:"name" binding:"required" example:"test-service"`
	Level   string `json:"level" binding:"required" example:"info"`
	Message string `json:"message" binding:"required" example:"This is a test log message"`
}

// Response models
type WriteLogResponse struct {
	Error   bool   `json:"error" example:"false"`
	Message string `json:"message" example:"Log written successfully"`
}

type ReadLogsResponse struct {
	Error bool            `json:"error" example:"false"`
	Logs  []data.LogEntry `json:"logs"`
}

type ErrorResponse struct {
	Error   bool   `json:"error" example:"true"`
	Message string `json:"message" example:"Error message"`
}

// @Summary Write a log message
// @Description Write a log message to the logging service
// @Tags Log
// @version 1.0
// @produce application/json
// @Param request body WriteLogRequest true "Log message details"
// @Success 200 {object} WriteLogResponse
// @Failure 400 {object} ErrorResponse "Bad Request - Invalid request body"
// @Failure 500 {object} ErrorResponse "Internal Server Error - Failed to write log"
// @Router /v1/log [post]
func (app *Config) WriteLogViaHTTP(c *gin.Context) {

	var request WriteLogRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   true,
			Message: "Failed to parse request body",
		})
		return
	}

	logEntry := data.LogEntry{
		Name:    request.Name,
		Level:   request.Level,
		Message: request.Message,
	}

	err := app.Repo.Insert(logEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   true,
			Message: "Failed to write log entry",
		})
		return
	}

	c.JSON(http.StatusCreated, WriteLogResponse{
		Error:   false,
		Message: "Log written successfully",
	})
}

// @Summary Read all logs
// @Description Retrieve all log entries from the logging service
// @Tags Log
// @version 1.0
// @produce application/json
// @Success 200 {object} ReadLogsResponse
// @Failure 500 {object} ErrorResponse "Internal Server Error - Failed to retrieve log entries"
// @Router /v1/log [get]
func (app *Config) ReadAllLogs(c *gin.Context) {

	logEntries, err := app.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   true,
			Message: "Failed to retrieve log entries",
		})
		return
	}

	c.JSON(http.StatusOK, ReadLogsResponse{
		Error: false,
		Logs:  *logEntries,
	})
}
