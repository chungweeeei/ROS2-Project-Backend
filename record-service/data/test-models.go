package data

import (
	"time"

	"gorm.io/gorm"
)

func TestNew(dbPool *gorm.DB) Models {
	db = dbPool

	return Models{
		TradeRecord: &TradeRecordTest{},
	}
}

type TradeRecordTest struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Email       string    `json:"email" gorm:"unique;type:character varying(255);index"`
	StockNumber string    `json:"stock_number" gorm:"type:character varying(255);not null"`
	StockName   string    `json:"stock_name" gorm:"type:character varying(255);not null"`
	Side        string    `json:"side" gorm:"type:text;not null"`
	EntryPrice  float64   `json:"entry_price" gorm:"type:float;not null"`
	ExitPrice   float64   `json:"exit_price" gorm:"type:float;not null"`
	Quantity    int       `json:"quantity" gorm:"type:integer;not null"`
	EntryTime   time.Time `json:"entry_time" gorm:"type:timestamp without time zone"`
	ExitTime    time.Time `json:"exit_time" gorm:"type:timestamp without time zone"`
	Notes       string    `json:"notes" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp without time zone"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp without time zone"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"type:timestamp without time zone"`
}

func (tr *TradeRecordTest) GetAll() ([]TradeRecord, error) {

	records := []TradeRecord{}

	testEntryTime, _ := time.Parse("2006-01-02 15:04:05", "2025-08-18 00:00:00")
	testExitTime, _ := time.Parse("2006-01-02 15:04:05", "2025-08-19 00:00:00")

	record := TradeRecord{
		ID:          1,
		Email:       "test@test.com",
		StockNumber: "2330",
		StockName:   "台積電",
		Side:        "buy",
		EntryPrice:  1185.0,
		ExitPrice:   1200.0,
		Quantity:    1000,
		EntryTime:   testEntryTime,
		ExitTime:    testExitTime,
		Notes:       "Test record",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	records = append(records, record)

	return records, nil
}

func (tr *TradeRecordTest) Insert(record TradeRecord) (int, error) {
	return 2, nil
}

func (tr *TradeRecordTest) Delete(id int) error {
	return nil
}
