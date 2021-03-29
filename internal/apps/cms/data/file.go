package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/whole/internal/apps/cms/biz"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewCategoryRepo .
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper("category_repo", logger),
	}
}

func (ar *categoryRepo) ListCategory(ctx context.Context) ([]*biz.Category, error) {
	ps, err := ar.data.db.Category.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Category, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Category{
			ID: p.ID,
		})
	}
	return rv, nil
}

func (ar *categoryRepo) GetCategory(ctx context.Context, id int64) (*biz.Category, error) {
	p, err := ar.data.db.Category.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Category{
		ID: p.ID,
	}, nil
}

func (ar *categoryRepo) CreateCategory(ctx context.Context, category *biz.Category) error {
	_, err := ar.data.db.Category.
		Create().
		Save(ctx)
	return err
}

func (ar *categoryRepo) UpdateCategory(ctx context.Context, id int64, category *biz.Category) error {
	p, err := ar.data.db.Category.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		Save(ctx)
	return err
}

func (ar *categoryRepo) DeleteCategory(ctx context.Context, id int64) error {
	return ar.data.db.Category.DeleteOneID(id).Exec(ctx)
}
