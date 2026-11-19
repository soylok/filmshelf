package posters

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func (s *PosterStorage) Delete(filename string) error {
	path := filepath.Join(s.postersDir, filename)
	if err := os.Remove(path); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("remove poster file: %w", err)
	}
	return nil
}
