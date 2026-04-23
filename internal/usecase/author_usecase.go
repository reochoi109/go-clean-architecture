package usecase

import (
	"context"
	"go-clean-architecture/domain"
)

type authorUsecase struct {
	authorRepo  domain.AuthorRepository
	articleRepo domain.ArticleRepository
}

var _ domain.AuthorUsecase = (*authorUsecase)(nil)

func NewAuthorUsecase(aRepo domain.AuthorRepository, arRepo domain.ArticleRepository) domain.AuthorUsecase {
	return &authorUsecase{
		authorRepo:  aRepo,
		articleRepo: arRepo,
	}
}

func (u *authorUsecase) GetAuthorWithArticles(ctx context.Context, id string) (domain.Author, []domain.Article, error) {
	author, err := u.authorRepo.GetByID(ctx, id)
	if err != nil {
		return domain.Author{}, nil, err
	}

	articles, err := u.articleRepo.GetByAuthorID(ctx, author.ID)
	return author, articles, err
}
