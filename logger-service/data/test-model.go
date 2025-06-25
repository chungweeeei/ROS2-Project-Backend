package data

import (
	"time"

	"gorm.io/gorm"
)

type PostgresTestRepository struct {
	Conn *gorm.DB
}

func NewPostgresTestRepository(dbPool *gorm.DB) (*PostgresTestRepository, error) {
	return &PostgresTestRepository{
		Conn: dbPool,
	}, nil
}

func (p *PostgresTestRepository) Insert(logEntry LogEntry) error {
	// This is a mock implementation for testing purposes.
	// In a real application, you would insert the log entry into the database.
	return nil
}

func (p *PostgresTestRepository) GetAll() (*[]LogEntry, error) {

	test1_entry := LogEntry{
		ID:        1,
		Name:      "Test Log",
		Level:     "info",
		Message:   "This is the first test log message",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	test2_entry := LogEntry{
		ID:        2,
		Name:      "Test Log",
		Level:     "info",
		Message:   "This is the second test log message",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	logs := []LogEntry{test1_entry, test2_entry}

	return &logs, nil
}
