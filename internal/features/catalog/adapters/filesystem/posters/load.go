package posters

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *PosterStorage) Load(src io.Reader, ext string) (string, error) {
	filename := uuid.NewString() + ext
	path := filepath.Join(s.postersDir, filename)

	dst, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return "", fmt.Errorf("create poster file: %w", err)
	}
	defer dst.Close()

	_, copyErr := io.Copy(dst, src)
	closeErr := dst.Close()
	if copyErr != nil {
		err = fmt.Errorf("write poster file: %w", copyErr)
	}
	if closeErr != nil {
		err = errors.Join(err, fmt.Errorf("close poster file: %w", closeErr))
	}
	if err != nil {
		if rmErr := os.Remove(path); rmErr != nil {
			err = errors.Join(err, fmt.Errorf("remove partial poster file: %w", rmErr))
		}
		return "", err
	}
	return filename, nil
}
