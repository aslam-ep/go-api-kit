package auth

import (
	"context"
	"database/sql"
)

type AuthRepository interface {
	Save(ctx context.Context, refreshToken *RefreshToken) (*RefreshToken, error)
	Delete(ctx context.Context, refreshToken *RefreshToken) error
	FindByToken(ctx context.Context, token string) (*RefreshToken, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Save(ctx context.Context, refreshToken *RefreshToken) (*RefreshToken, error) {
	var refreshTokenID int
	insertQuery := `INSERT INTO refresh_tokens(user_id, token, expires_at) VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRowContext(ctx, insertQuery,
		refreshToken.UserID,
		refreshToken.Token,
		refreshToken.ExpiresAt,
	).Scan(&refreshTokenID)

	if err != nil {
		return nil, err
	}

	refreshToken.ID = int64(refreshTokenID)

	return refreshToken, nil
}

func (r *authRepository) Delete(ctx context.Context, refreshToken *RefreshToken) error {
	deleteQuery := `DELETE FROM refresh_tokens WHERE id = $1`

	_, err := r.db.ExecContext(ctx, deleteQuery, refreshToken.ID)

	return err
}

func (r *authRepository) FindByToken(ctx context.Context, token string) (*RefreshToken, error) {
	var refeshToken RefreshToken
	selectQueryByID := `SELECT * FROM refresh_tokens WHERE token = $1 AND expires_at > CURRENT_TIMESTAMP`

	err := r.db.QueryRowContext(ctx, selectQueryByID, token).Scan(
		&refeshToken.ID,
		&refeshToken.UserID,
		&refeshToken.Token,
		&refeshToken.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return &refeshToken, nil
}
