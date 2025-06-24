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

func (p *PostgresTestRepository) Insert(user User) (int, error) {
	return 1, nil
}

func (p *PostgresTestRepository) GetByEmail(email string) (*User, error) {

	user := User{
		ID:        1,
		Email:     "test@test.com",
		Username:  "testuser",
		Password:  "$2a$14$.tw7zquwxZW2mD/lz24C4OUc92KPZET4Xqoxn7UdYGgrd7Rd6.z6G",
		Role:      "guest",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}
