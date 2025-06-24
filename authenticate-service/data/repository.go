package data

type Repository interface {
	Insert(user User) (int, error)
	GetByEmail(email string) (*User, error)
}
