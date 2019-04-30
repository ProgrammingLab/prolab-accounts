package store

// ImageStore provides images
type ImageStore interface {
	CreateImage(image []byte) (filename string, err error)
	DeleteImage(filename string) error
}
