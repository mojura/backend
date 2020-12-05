package backend

// Initializer represents a Backend initializer
type Initializer interface {
	New(filename string) (Backend, error)
}
