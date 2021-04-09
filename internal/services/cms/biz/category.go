package biz

import (
	"context"
	"time"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/realotz/whole/api/cms/v1"
)

type Category struct {
	Id         int64
	Name       string
	Pid        int64
	Icon       string
	Desc       string
	UpdateTime time.Time
	CreateTime time.Time
}

type CategoryRepo interface {
	ListCategory(ctx context.Context, op *pb.ListCategoryRequest) ([]*Category, error)
	GetCategory(ctx context.Context, id int64) (*Category, error)
	CreateCategory(ctx context.Context, mod *Category) (*Category, error)
	UpdateCategory(ctx context.Context, id int64, mod *Category) (*Category, error)
	DeleteCategory(ctx context.Context, ids []int64) error
}
type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, log: log.NewHelper("biz/category", logger)}
}

//查询
func (uc *CategoryUsecase) Get(ctx context.Context, id int64) (p *Category, err error) {
	return uc.repo.GetCategory(ctx, id)
}

//列表
func (uc *CategoryUsecase) List(ctx context.Context, op *pb.ListCategoryRequest) ([]*Category, error) {
	return uc.repo.ListCategory(ctx, op)
}

//创建
func (uc *CategoryUsecase) Create(ctx context.Context, mod *Category) (p *Category, err error) {
	return uc.repo.CreateCategory(ctx, mod)
}

//创建
func (uc *CategoryUsecase) Update(ctx context.Context, mod *Category) (p *Category, err error) {
	return uc.repo.UpdateCategory(ctx, mod.Id, mod)
}

//删除
func (uc *CategoryUsecase) Delete(ctx context.Context, ids []int64) error {
	return uc.repo.DeleteCategory(ctx, ids)
}
