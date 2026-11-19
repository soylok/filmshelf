package media

import (
	"fmt"
	"os"
)

func EnsurePostersDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("ensure posters directory: %w", err)
	}
	return nil
}
