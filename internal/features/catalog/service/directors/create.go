package directors

import (
	"context"
	"strings"

	"github.com/soylok/filmshelf/internal/core/apperrors"
	"github.com/soylok/filmshelf/internal/core/domain"
)

func (s *DirectorService) Create(ctx context.Context, director domain.Director) (domain.Director, error) {
	director.FirstName = strings.Join(strings.Fields(director.FirstName), " ")
	director.LastName = strings.Join(strings.Fields(director.LastName), " ")
	if director.MiddleName != nil {
		normalized := strings.Join(strings.Fields(*director.MiddleName), " ")
		director.MiddleName = &normalized
	}

	if err := director.Validate(); err != nil {
		return domain.Director{}, &apperrors.Error{
			Type:    apperrors.ErrInvalidArgument,
			Message: err.Error(),
			Cause:   err,
		}
	}
	return s.repo.Create(ctx, director)
}
