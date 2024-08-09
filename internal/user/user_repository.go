package user

import (
	"context"
	"database/sql"
	"time"
)

// Repository interface for the user repository
type Repository interface {
	// Create stores a new user and returns the created user.
	Create(ctx context.Context, user *User) (*User, error)

	// GetByEmail find and returns the user by user email
	GetByEmail(ctx context.Context, email string) (*User, error)

	// GetByID find and returns the user, by user id
	GetByID(ctx context.Context, id int) (*User, error)

	// Update update user by user id and returns the updated user.
	Update(ctx context.Context, user *User) (*User, error)

	// ChangePassword update the user password by the user id
	ChangePassword(ctx context.Context, userID int, password string) error

	// Delete delete the given user based on user id
	Delete(ctx context.Context, userID int) error
}

type repository struct {
	db *sql.DB
}

// NewRepository initialize and return the Repository
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *User) (*User, error) {
	var userID int
	var (
		createdAt time.Time
		updatedAt time.Time
	)
	insertQuery := `INSERT INTO users(name, email, phone, role, password) VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`

	err := r.db.QueryRowContext(ctx, insertQuery,
		user.Name,
		user.Email,
		user.Phone,
		user.Role,
		user.Password,
	).Scan(&userID, &createdAt, &updatedAt)

	if err != nil {
		return nil, err
	}

	// Adding db generated values to user
	user.ID = int64(userID)
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt

	return user, nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	selectQueryByEmail := `SELECT id, name, email, phone, role, password, created_at, updated_at role FROM users WHERE email = $1 AND is_deleted = false`

	err := r.db.QueryRowContext(ctx, selectQueryByEmail, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Role,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetByID(ctx context.Context, id int) (*User, error) {
	var user User
	selectQueryByID := `SELECT id, name, email, phone, role, password, created_at, updated_at FROM users WHERE id = $1 AND is_deleted = false`

	err := r.db.QueryRowContext(ctx, selectQueryByID, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Role,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Update(ctx context.Context, user *User) (*User, error) {
	user.UpdatedAt = time.Now()
	updateQuery := `UPDATE users SET name = $1, phone = $2, role = $3, updated_at = $4 WHERE id = $5`

	_, err := r.db.ExecContext(ctx, updateQuery,
		user.Name,
		user.Phone,
		user.Role,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) ChangePassword(ctx context.Context, userID int, password string) error {
	passwordUpdateQuery := `UPDATE users SET password = $1 WHERE id = $2`

	_, err := r.db.ExecContext(ctx, passwordUpdateQuery, password, userID)

	return err
}

func (r *repository) Delete(ctx context.Context, userID int) error {
	deleteQuery := `UPDATE users SET is_deleted = true WHERE id = $1`

	_, err := r.db.ExecContext(ctx, deleteQuery, userID)

	return err
}
