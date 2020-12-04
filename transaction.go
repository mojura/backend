package backend

// Transaction represents a database transaction
type Transaction interface {
	GetBucket(key []byte) Bucket
	GetOrCreateBucket(key []byte) (Bucket, error)
}
