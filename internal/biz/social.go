package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
}
type CommentRepo interface {
}
type TagRepo interface {
}

type SocialUsecase struct {
	article ArticleRepo
	comment CommentRepo
	tag     TagRepo

	log *log.Helper
}

func NewSocialUsecase(article ArticleRepo,
	comment CommentRepo,
	tag TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{article: article, comment: comment, tag: tag , log: log.NewHelper(logger)}
}

func (uc *SocialUsecase) CreateArticle(ctx context.Context) error {
	return nil
}

func (uc *SocialUsecase) UpdateArticle(ctx context.Context) error {
	return nil
}
