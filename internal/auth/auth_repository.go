package auth

import (
	"context"
	"database/sql"
)

// Repository interface for auth repository
type Repository interface {
	// Save stores a new refresh token in the data store.
	Save(ctx context.Context, refreshToken *RefreshToken) (*RefreshToken, error)

	// Delete removes a refresh token from the data store.
	Delete(ctx context.Context, refreshTokenID int) error

	// FindByToken retrieves a refresh token by its token string from the data store.
	FindByToken(ctx context.Context, token string) (*RefreshToken, error)
}

type repository struct {
	db *sql.DB
}

// NewRepository initialize and returns auth repository
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, refreshToken *RefreshToken) (*RefreshToken, error) {
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

func (r *repository) Delete(ctx context.Context, refreshTokenID int) error {
	deleteQuery := `DELETE FROM refresh_tokens WHERE id = $1`

	_, err := r.db.ExecContext(ctx, deleteQuery, refreshTokenID)

	return err
}

func (r *repository) FindByToken(ctx context.Context, token string) (*RefreshToken, error) {
	var refreshToken RefreshToken
	selectQueryByID := `SELECT * FROM refresh_tokens WHERE token = $1 AND expires_at > CURRENT_TIMESTAMP`

	err := r.db.QueryRowContext(ctx, selectQueryByID, token).Scan(
		&refreshToken.ID,
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}
