package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/murashi19/koda-b8-backend1/internal/models"
)

type UserRepo struct {
	// data *[]models.User
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, data *models.User) (*models.User, error) {

	sql := `
		INSERT INTO users (
			email,
			password,
			username,
			phone,
			picture
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`

	user, err := oneRow[models.User](ctx, r.db, sql, data.Email, data.Password, data.Username, data.Phone, data.Picture)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, fmt.Errorf("email already registered")
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) GetById(ctx context.Context, id int64) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = $1;`

	return oneRow[models.User](ctx, r.db, query, id)
}

func (r *UserRepo) UpdateUser(ctx context.Context, id int64, data *models.UpdateUserRequest) (*models.User, error) {
	query := `UPDATE users SET 
	email = COALESCE($1, email),
	username = COALESCE($2, username),
	phone = COALESCE($3, phone),
	updated_at = NOW()
	WHERE id = $4
	RETURNING *
	`
	return oneRow[models.User](ctx, r.db, query, data.Email, data.Username, data.Phone, id)
}

func (r *UserRepo) DeleteUser(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	rows, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if rows.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
	SELECT
		id,
		email,
		password,
		username,
		phone,
		created_at,
		updated_at,
		picture
	FROM users
	WHERE email = $1
	`

	return oneRow[models.User](ctx, r.db, query, email)
}

func (r *UserRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	sql := `SELECT * FROM users
			ORDER BY id`
	return rows[models.User](ctx, r.db, sql)
}

func (r *UserRepo) Upload(ctx context.Context, id int64, picture string) (*models.User, error) {
	sql := `UPDATE users SET 
	picture = $1,
	updated_at = NOW()
	WHERE id = $2
	RETURNING *
	`
	return oneRow[models.User](ctx, r.db, sql, picture, id)
}
