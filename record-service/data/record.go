package data

import (
	"errors"
	"time"
)

type TradeRecord struct {
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

func (tr *TradeRecord) GetAll() ([]TradeRecord, error) {

	records := []TradeRecord{}

	result := db.Find(&records)
	if result.Error != nil {
		return nil, errors.New("failed to retrieve trade records")
	}

	return records, nil
}

func (tr *TradeRecord) Insert(record TradeRecord) (int, error) {

	result := db.Create(&record)
	if result.Error != nil {
		return 0, errors.New("failed to insert trade record")
	}

	return record.ID, nil
}

func (tr *TradeRecord) Delete(id int) error {

	result := db.Delete(&TradeRecord{}, id)
	if result.Error != nil {
		return errors.New("failed to delete trade record")
	}
	return nil

}
