package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/whole/internal/apps/systems/biz"
)

type fileRepo struct {
	data *Data
	log  *log.Helper
}

// NewFileRepo .
func NewFileRepo(data *Data, logger log.Logger) biz.FileRepo {
	return &fileRepo{
		data: data,
		log:  log.NewHelper("file_repo", logger),
	}
}

func (ar *fileRepo) ListFile(ctx context.Context) ([]*biz.File, error) {
	ps, err := ar.data.db.File.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.File, 0)
	for _, p := range ps {
		rv = append(rv, &biz.File{})
	}
	return rv, nil
}

func (ar *fileRepo) GetFile(ctx context.Context, id int64) (*biz.File, error) {
	p, err := ar.data.db.File.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.File{}, nil
}

func (ar *fileRepo) CreateFile(ctx context.Context, file *biz.File) error {
	_, err := ar.data.db.File.
		Create().
		Save(ctx)
	return err
}

func (ar *fileRepo) UpdateFile(ctx context.Context, id int64, file *biz.File) error {
	p, err := ar.data.db.File.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		Save(ctx)
	return err
}

func (ar *fileRepo) DeleteFile(ctx context.Context, id int64) error {
	return ar.data.db.File.DeleteOneID(id).Exec(ctx)
}
