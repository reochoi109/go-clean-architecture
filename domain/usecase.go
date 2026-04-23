package domain

import "context"

type AuthorUsecase interface {
	GetAuthorWithArticles(ctx context.Context, id string) (Author, []Article, error)
}
