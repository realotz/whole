package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Category struct {
	ID         int64     `json:"id,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}

type CategoryRepo interface {
	ListCategory(ctx context.Context) ([]*Category, error)
	GetCategory(ctx context.Context, id int64) (*Category, error)
	CreateCategory(ctx context.Context, article *Category) error
	UpdateCategory(ctx context.Context, id int64, article *Category) error
	DeleteCategory(ctx context.Context, id int64) error
}

type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, log: log.NewHelper("biz/category", logger)}
}

func (uc *CategoryUsecase) List(ctx context.Context) (ps []*Category, err error) {
	return uc.repo.ListCategory(ctx)
}

func (uc *CategoryUsecase) Get(ctx context.Context, id int64) (p *Category, err error) {
	return uc.repo.GetCategory(ctx, id)
}

func (uc *CategoryUsecase) Create(ctx context.Context, article *Category) error {
	return uc.repo.CreateCategory(ctx, article)
}

func (uc *CategoryUsecase) Update(ctx context.Context, id int64, article *Category) error {
	return uc.repo.UpdateCategory(ctx, id, article)
}

func (uc *CategoryUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteCategory(ctx, id)
}
