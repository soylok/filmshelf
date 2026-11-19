package movies

import (
	"context"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"path/filepath"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	outbound "github.com/soylok/filmshelf/internal/features/catalog/ports/out"
)

// maxPosterSize must be below the transport's total multipart body
// limit maxBodySize.
const maxPosterSize int64 = 9 << 20

var allowedPosterExtensions = map[string]string{
	".jpg":  "jpeg",
	".jpeg": "jpeg",
	".png":  "png",
}

func (s *MovieService) UpdatePoster(
	ctx context.Context,
	id int64,
	file io.ReadSeeker,
	filename string,
	size int64,
) (bool, error) {
	var replaced bool
	oldFilename, err := s.repo.GetPosterFilename(ctx, id)
	if err != nil {
		if errors.Is(err, outbound.ErrNotFound) {
			return replaced, &apperrors.Error{
				Type:    apperrors.ErrNotFound,
				Message: "film not found",
				Cause:   err,
			}
		}
		return replaced, err
	}
	if oldFilename != nil {
		replaced = true
	}

	ext := filepath.Ext(filename)
	expectedFormat, ok := allowedPosterExtensions[ext]
	if !ok {
		return replaced, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: "unsupported poster file extension",
		}
	}

	if size > maxPosterSize {
		return replaced, &apperrors.Error{
			Type:    apperrors.ErrPayloadTooLarge,
			Message: "poster file too large",
		}
	}

	cfg, format, err := image.DecodeConfig(file)
	if err != nil {
		return replaced, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: "invalid poster image",
			Cause:   err,
		}
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return replaced, fmt.Errorf("rewind poster file: %w", err)
	}
	if format != expectedFormat {
		return replaced, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: "poster image format does not match file extension",
		}
	}
	ratio := float64(cfg.Width) / float64(cfg.Height)
	if ratio < 0.4 || ratio > 1.2 {
		return replaced, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: "poster width-to-height ratio must be between 0.4 and 1.2",
		}
	}

	newFilename, err := s.posterStorage.Load(file, ext)
	if err != nil {
		return replaced, fmt.Errorf("load movie poster: %w", err)
	}

	err = s.repo.UpdatePoster(ctx, id, oldFilename, newFilename)
	if err != nil {
		rmErr := s.posterStorage.Delete(newFilename)
		if rmErr != nil {
			err = errors.Join(err, fmt.Errorf("delete new movie poster: %w", err))
		}
		if errors.Is(err, outbound.ErrConcurrentModification) {
			return replaced, &apperrors.Error{
				Type:    apperrors.ErrConflict,
				Message: "movie or poster was modified or deleted; reload and try again",
				Cause:   err,
			}
		}
		return replaced, err
	}

	if oldFilename != nil {
		if err := s.posterStorage.Delete(*oldFilename); err != nil {
			return replaced, fmt.Errorf("delete old movie poster: %w", err)
		}
	}
	return replaced, nil
}
