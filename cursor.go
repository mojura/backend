package backend

// Cursor represents a cursor
type Cursor interface {
	Seek(seekTo []byte) (key, value []byte)
	First() (key, value []byte)
	Next() (key, value []byte)
	Prev() (key, value []byte)
	Last() (key, value []byte)
}
