package backend

// Bucket represents a bucket
type Bucket interface {
	Transaction

	Get(key []byte) (value []byte)
	Put(key, value []byte) error
	Delete(key []byte) error
	Cursor() Cursor
	ForEach(func(key, value []byte) error) error
}
