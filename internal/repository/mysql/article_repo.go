package mysql

import (
	"context"
	"database/sql"
	"go-clean-architecture/domain"
)

type ArticleRepository struct{ db *sql.DB }

var _ domain.ArticleRepository = &ArticleRepository{}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetByAuthorID(ctx context.Context, authorID string) ([]domain.Article, error) {
	return []domain.Article{}, nil
}

func (r *ArticleRepository) Store(ctx context.Context, a *domain.Article) error {
	return nil
}
