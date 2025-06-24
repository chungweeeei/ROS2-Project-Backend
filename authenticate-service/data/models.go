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
	err := dbEngine.AutoMigrate(&User{})
	if err != nil {
		return &PostgresRepository{}, errors.New("failed to auto migrate user model")
	}

	return &PostgresRepository{
		Conn: dbEngine,
	}, nil
}

// // New function for initializing the Models struct with a User model
// func New(dbEngine *gorm.DB) (Models, error) {

// 	// assign to the global db variable
// 	db = dbEngine

// 	// // Do auto migration
// 	err := dbEngine.AutoMigrate(&User{})
// 	if err != nil {
// 		return Models{}, errors.New("failed to auto migrate user model")
// 	}

// 	return Models{User: User{}}, nil
// }

// type Models struct {
// 	User User
// }

// gorm.Model definition
type User struct {
	ID        int       `json:"id" gorm:"primaryKey;not null"`
	Email     string    `json:"email" gorm:"unique;required;index;not null"`
	Username  string    `json:"username" gorm:"required;not null"`
	Password  string    `json:"password" gorm:"required;not null"`
	Role      string    `json:"role" gorm:"default:'user';not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *PostgresRepository) Insert(user User) (int, error) {

	result := db.Create(&user)

	if result.Error != nil {
		return 0, errors.New("failed to insert user into database")
	}

	return user.ID, nil
}

func (u *PostgresRepository) GetByEmail(email string) (*User, error) {

	var user User
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &User{}, errors.New("user not found")
		}
		return &User{}, errors.New("failed to retrieve user by email")
	}

	return &user, nil
}
