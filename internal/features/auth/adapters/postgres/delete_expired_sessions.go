package postgres

import (
	"context"
	"fmt"
)

func (repo *AuthRepository) DeleteExpiredSessions(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `delete
    from filmshelf.sessions
    where expires_at <= now()`

	if _, err := repo.querier.Exec(ctx, query); err != nil {
		return fmt.Errorf("delete expired sessions: %w", err)
	}
	return nil
}