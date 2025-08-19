package main

import (
	"logger-service/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WriteLogRequest struct {
	Message string `json:"message" binding:"required" example:"This is a test log message"`
	Level   string `json:"level" binding:"required" example:"info"`
}

type LogsResponse struct {
	Logs []data.Log `json:"logs"`
}

func (app *Config) ReadAllLogs(c *gin.Context) {

	logs, err := app.Models.Log.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve logs",
		})
		return
	}

	c.JSON(http.StatusOK, LogsResponse{
		Logs: logs,
	})
}

func (app *Config) WriteLog(c *gin.Context) {

	var request WriteLogRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse request body",
		})
		return
	}

	err := app.Models.Log.Insert(data.Log{
		Message: request.Message,
		Level:   request.Level,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to write log to database",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Log written successfully",
	})
}
