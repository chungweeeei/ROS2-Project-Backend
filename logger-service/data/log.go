package data

import (
	"errors"
	"time"
)

type Log struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Message   string    `json:"message" gorm:"type:text;not null"`
	Level     string    `json:"level" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp without time zone"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp without time zone"`
}

func (l *Log) GetAll() ([]Log, error) {

	var logs []Log
	result := db.Find(&logs)
	if result.Error != nil {
		return nil, errors.New("failed to retrieve logs")
	}
	return logs, nil
}

func (l *Log) Insert(log Log) error {

	result := db.Create(&log)
	if result.Error != nil {
		return errors.New("failed to insert log")
	}
	return nil
}

func (l *Log) Delete(id int) error {

	result := db.Delete(&Log{}, id)
	if result.Error != nil {
		return errors.New("failed to delete log")
	}
	return nil
}
