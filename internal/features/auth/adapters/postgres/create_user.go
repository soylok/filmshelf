package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/soylok/filmshelf/internal/core/domain"
	outbound "github.com/soylok/filmshelf/internal/features/auth/ports/out"
)

func (repo *AuthRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, repo.opTimeout)
	defer cancel()

	query := `insert into filmshelf.users(
		first_name,
		second_name,
		email,
		password_hash
	) values
		($1, $2, $3, $4)
	returning id, first_name, second_name, email, email_verified, role, version`
	var userModel UserModel
	row := repo.querier.QueryRow(ctx, query, user.FirstName, user.SecondName, user.Email, user.PasswordHash)
	err := row.Scan(
		&userModel.ID,
		&userModel.FirstName,
		&userModel.SecondName,
		&userModel.Email,
		&userModel.EmailVerified,
		&userModel.Role,
		&userModel.Version,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.User{}, outbound.ErrAlreadyExists
		}
		return domain.User{}, fmt.Errorf("scan user: %w", err)
	}
	return userModel.ToDomain(), nil
}
