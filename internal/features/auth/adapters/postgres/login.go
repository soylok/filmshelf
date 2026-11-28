package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/soylok/filmshelf/internal/core/domain"
	outbound "github.com/soylok/filmshelf/internal/features/auth/ports/out"
)

func (repo *AuthRepository) Login(ctx context.Context, email string) (domain.User, string, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `select
		id,
		first_name,
		second_name,
		email,
		email_verified,
		password_hash,
		role,
		version
	from filmshelf.users
	where email = $1`
	var userModel UserModel
	var passwordHash string
	row := repo.querier.QueryRow(ctx, query, email)
	err := row.Scan(
		&userModel.ID,
		&userModel.FirstName,
		&userModel.SecondName,
		&userModel.Email,
		&userModel.EmailVerified,
		&passwordHash,
		&userModel.Role,
		&userModel.Version,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, "", outbound.ErrNotFound
		}
		return domain.User{}, "", fmt.Errorf("scan user: %w", err)
	}
	return userModel.ToDomain(), passwordHash, nil
}
