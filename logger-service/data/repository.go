package data

type Repository interface {
	Insert(logEntry LogEntry) error
	GetAll() ([]LogEntry, error)
}
