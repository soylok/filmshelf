package outbound

import "io"

type PosterStorage interface {
	Delete(filename string) error
	Load(src io.Reader, ext string) (string, error)
}
