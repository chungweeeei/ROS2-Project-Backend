package data

import (
	"time"

	"gorm.io/gorm"
)

func TestNew(dbPool *gorm.DB) Models {
	db = dbPool

	return Models{
		Log: &LogTest{},
	}
}

type LogTest struct {
	ID        int       `json:"id" gorm:"primaryKey;not null"`
	Message   string    `json:"message" gorm:"type:text;not null"`
	Level     string    `json:"level" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp without time zone"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp without time zone"`
}

func (l *LogTest) GetAll() ([]Log, error) {

	logs := []Log{}

	log := Log{
		ID:        1,
		Message:   "Test log message",
		Level:     "info",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	logs = append(logs, log)

	return logs, nil
}

func (l *LogTest) Insert(log Log) error {
	return nil
}

func (l *LogTest) Delete(id int) error {
	return nil
}
