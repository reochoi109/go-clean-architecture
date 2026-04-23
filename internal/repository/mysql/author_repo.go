package mysql

import (
	"context"
	"database/sql"
	"go-clean-architecture/domain"
)

type AuthorRepository struct {
	db *sql.DB
}

var _ domain.AuthorRepository = &AuthorRepository{}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) GetByID(ctx context.Context, id string) (domain.Author, error) {
	query := `SELECT id, email, name, created_at, updated_at FROM author WHERE id=?`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return domain.Author{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	res := domain.Author{}
	if err = row.Scan(&res.ID, &res.Email, &res.Name, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return domain.Author{}, err
	}

	return res, nil
}

func (r *AuthorRepository) Store(ctx context.Context, a *domain.Author) error {
	query := `INSERT INTO author (id, email, name, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, a.ID, a.Email, a.Name, a.CreatedAt, a.UpdatedAt)
	return err
}

func (r *AuthorRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM author WHERE id = ?`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
