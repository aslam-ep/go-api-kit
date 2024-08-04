package user

import (
	"context"
	"database/sql"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	ResetPassword(ctx context.Context, user *User, password string) error
	Delete(ctx context.Context, user *User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *User) (*User, error) {
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

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	selectQueryByEmail := `SELECT id, name, email, phone, password, created_at, updated_at role FROM users WHERE email = $1 AND is_deleted = false`

	err := r.db.QueryRowContext(ctx, selectQueryByEmail, email).Scan(
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

func (r *userRepository) GetByID(ctx context.Context, id int) (*User, error) {
	var user User
	selectQueryByID := `SELECT id, name, email, phone, password, created_at, updated_at FROM users WHERE id = $1 AND is_deleted = false`

	err := r.db.QueryRowContext(ctx, selectQueryByID, id).Scan(
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

func (r *userRepository) Update(ctx context.Context, user *User) (*User, error) {
	user.UpdatedAt = time.Now()
	updateQuery := `UPDATE users SET name = $1, phone = $2, role = $3, updated_at = $4 WHERE id = $5`

	_, err := r.db.ExecContext(ctx, updateQuery,
		&user.Name,
		&user.Phone,
		&user.Role,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) ResetPassword(ctx context.Context, user *User, password string) error {
	passwordUpdateQueryr := `UPDATE users SET passsword = $1 WHERE id = $2`

	_, err := r.db.ExecContext(ctx, passwordUpdateQueryr, password, user.ID)

	return err
}

func (r *userRepository) Delete(ctx context.Context, user *User) error {
	deleteQuery := `UPDATE users SET is_deleted = true WHERE id = $1`

	_, err := r.db.ExecContext(ctx, deleteQuery, user.ID)

	return err
}
