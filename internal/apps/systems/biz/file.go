package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type File struct {
	ID         int64     `json:"id,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}

type FileRepo interface {
	ListFile(ctx context.Context) ([]*File, error)
	GetFile(ctx context.Context, id int64) (*File, error)
	CreateFile(ctx context.Context, article *File) error
	UpdateFile(ctx context.Context, id int64, article *File) error
	DeleteFile(ctx context.Context, id int64) error
}

type FileUsecase struct {
	repo FileRepo
	log  *log.Helper
}

func NewFileUsecase(repo FileRepo, logger log.Logger) *FileUsecase {
	return &FileUsecase{repo: repo, log: log.NewHelper("biz/file", logger)}
}

func (uc *FileUsecase) List(ctx context.Context) (ps []*File, err error) {
	return uc.repo.ListFile(ctx)
}

func (uc *FileUsecase) Get(ctx context.Context, id int64) (p *File, err error) {
	return uc.repo.GetFile(ctx, id)
}

func (uc *FileUsecase) Create(ctx context.Context, article *File) error {
	return uc.repo.CreateFile(ctx, article)
}

func (uc *FileUsecase) Update(ctx context.Context, id int64, article *File) error {
	return uc.repo.UpdateFile(ctx, id, article)
}

func (uc *FileUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteFile(ctx, id)
}
