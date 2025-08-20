package main

import (
	"net/http"
	"record-service/data"

	"github.com/gin-gonic/gin"
)

type WriteRecordRequest struct {
	Side        string  `json:"side" binding:"required" example:"buy"`
	StockNumber string  `json:"stock_number" binding:"required" example:"2330"`
	StockName   string  `json:"stock_name" binding:"required" example:"台積電"`
	EntryPrice  float64 `json:"entry_price" binding:"required" example:"100.50"`
	ExitPrice   float64 `json:"exit_price" binding:"required" example:"105.00"`
	Quantity    int     `json:"quantity" binding:"required" example:"10"`
	EntryTime   string  `json:"entry_time" binding:"required" example:"2025-08-18"`
	ExitTime    string  `json:"exit_time" binding:"required" example:"2025-08-19"`
	Notes       string  `json:"notes" binding:"required" example:"Test trade"`
}

type RecordsResponse struct {
	Records []data.TradeRecord `json:"records"`
}

func (app *Config) FetchAllRecords(c *gin.Context) {

	records, err := app.Models.TradeRecord.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve trade records",
		})
		return
	}

	c.JSON(http.StatusOK, RecordsResponse{
		Records: records,
	})
}

func (app *Config) WriteRecord(c *gin.Context) {

	var request WriteRecordRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse request body",
		})
		return
	}

	_, err := app.Models.TradeRecord.Insert(data.TradeRecord{
		Email:       c.GetString("Email"),
		StockNumber: request.StockNumber,
		StockName:   request.StockName,
		Side:        request.Side,
		EntryPrice:  request.EntryPrice,
		ExitPrice:   request.ExitPrice,
		Quantity:    request.Quantity,
		EntryTime:   convertStringToTime(request.EntryTime),
		ExitTime:    convertStringToTime(request.ExitTime),
		Notes:       request.Notes,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to write trade record to database",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Trade record written successfully",
	})
}
