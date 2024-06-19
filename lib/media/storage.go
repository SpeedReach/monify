package media

type Storage interface {
	// Store
	// Stores the image in the storage and returns the URL
	Store(path string, data []byte) error

	Delete(path string) error

	GetHost() string
	GetUrl(path string) string
}
