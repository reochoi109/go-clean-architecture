package domain

import "context"

type AuthorRepository interface {
	GetByID(ctx context.Context, id string) (Author, error)
	Store(ctx context.Context, a *Author) error
	Delete(ctx context.Context, id string) error
}

type ArticleRepository interface {
	GetByAuthorID(ctx context.Context, authorID string) ([]Article, error)
	Store(ctx context.Context, a *Article) error
}
