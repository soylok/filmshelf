package postgres

import (
	"context"
	"time"
)

func (repo *AuthRepository) CreateSession(ctx context.Context, tokenHash string, userID int64, expiresAt time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `insert into filmshelf.sessions(token_hash, user_id, expires_at)
	values
		($1, $2, $3)`
	_, err := repo.querier.Exec(ctx, query, tokenHash, userID, expiresAt)
	if err != nil {
		return err
	}
	return nil
}
