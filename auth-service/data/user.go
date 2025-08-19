package data

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `json:"id" gorm:"PrimaryKey;not null"`
	Email     string    `json:"email" gorm:"unique;type:character varying(255);index"`
	FirstName string    `json:"first_name" gorm:"type:character varying(255)"`
	LastName  string    `json:"last_name" gorm:"type:character varying(255)"`
	Password  string    `json:"password" gorm:"type:character varying(60)"`
	Active    int       `json:"active" gorm:"type:integer;default:0"`
	IsAdmin   int       `json:"is_admin" gorm:"type:integer;default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp without time zone"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp without time zone"`
}

func (u *User) PasswordMatches(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (u *User) GetAll() ([]*User, error) {

	users := []*User{}
	result := db.Find(&users)

	if result.Error != nil {
		return nil, errors.New("failed to retrieve users")
	}

	return users, nil
}

func (u *User) GetByEmail(email string) (*User, error) {

	user := User{}
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, errors.New("failed to retrieve user by email")
	}

	return &user, nil
}

func (u *User) GetOne(id int) (*User, error) {

	user := User{}
	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, errors.New("failed to retrieve user by id")
	}

	return &user, nil
}

func (u *User) Insert(user User) (int, error) {

	result := db.Create(&user)

	if result.Error != nil {
		return 0, errors.New("failed to insert user into database")
	}

	return user.ID, nil
}

func (u *User) Update(user User) error {

	result := db.Save(user)

	if result.Error != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func (u *User) Delete(id int) error {

	result := db.Delete(&User{}, id)

	if result.Error != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
