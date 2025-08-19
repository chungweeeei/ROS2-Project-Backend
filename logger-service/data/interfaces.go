package data

type LogInterface interface {
	GetAll() ([]Log, error)
	Insert(log Log) error
	Delete(id int) error
}
