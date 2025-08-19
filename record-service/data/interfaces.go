package data

type TradeRecordInterface interface {
	GetAll() ([]TradeRecord, error)
	Insert(record TradeRecord) (int, error)
	Delete(id int) error
}
