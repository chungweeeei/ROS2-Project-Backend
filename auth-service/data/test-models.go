package data

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func TestNew(dbPool *gorm.DB) Models {
	db = dbPool

	return Models{
		User: &UserTest{},
	}
}

type UserTest struct {
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

func (u *UserTest) PasswordMatches(password string, hashedPassword string) bool {
	return password != "wrong"
}

func (u *UserTest) GetAll() ([]*User, error) {

	users := []*User{}

	user := User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "Admin",
		Password:  "abc",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users = append(users, &user)

	return users, nil
}

func (u *UserTest) GetByEmail(email string) (*User, error) {

	if email == "test-wrong@example.com" {
		return nil, errors.New("failed to retrieve user by email")
	}

	user := User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "Admin",
		Password:  "abc",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

func (u *UserTest) GetOne(id int) (*User, error) {

	return u.GetByEmail("")
}

func (u *UserTest) Insert(user User) (int, error) {

	if user.Email == "test-wrong@example.com" {
		return 0, errors.New("failed to insert user")
	}

	return 2, nil
}

func (u *UserTest) Update(user User) error {
	return nil
}

func (u *UserTest) Delete(id int) error {
	return nil
}
