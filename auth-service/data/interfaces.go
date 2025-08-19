package data

type UserInterface interface {
	GetAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
	GetOne(id int) (*User, error)
	Insert(user User) (int, error)
	Update(user User) error
	Delete(id int) error
	PasswordMatches(password string, hashedPassword string) bool
}
