package backend

// Bucket represents a bucket
type Bucket interface {
	Get(key []byte) (value []byte)
	Put(key, value []byte) error
	Delete(key []byte) error
	Cursor() Cursor
	GetBucket(key []byte) Bucket
	GetOrCreateBucket(key []byte) (Bucket, error)
	DeleteBucket(key []byte) error
	ForEach(func(key, value []byte) error) error
}
