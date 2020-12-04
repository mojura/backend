package backend

// Backend represents a database
type Backend interface {
	Transaction(func(Transaction) error) error
	ReadTransaction(func(Transaction) error) error
	Close() error
}
