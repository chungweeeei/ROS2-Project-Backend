package data

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

type PostgresRepository struct {
	Conn *gorm.DB
}

func NewPostgresRepository(dbEngine *gorm.DB) (*PostgresRepository, error) {

	db = dbEngine

	// Do auto migration
	err := dbEngine.AutoMigrate(&LogEntry{})
	if err != nil {
		return &PostgresRepository{}, err
	}

	return &PostgresRepository{
		Conn: dbEngine,
	}, nil
}

// Define the LogEntry gorm model
type LogEntry struct {
	ID        int       `json:"id:"primaryKey;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Message   string    `json:"message" gorm:"not null"`
	Level     string    `json:"level" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *PostgresRepository) Insert(logEntry LogEntry) error {

	result := db.Create(&logEntry)
	if result.Error != nil {
		return errors.New("failed to insert log entry")
	}

	return nil
}

func (p *PostgresRepository) GetAll() ([]LogEntry, error) {

	var logEntries []LogEntry
	result := db.Find(&logEntries)
	if result.Error != nil {
		return nil, errors.New("failed to retrieve log entries")
	}

	return logEntries, nil
}
