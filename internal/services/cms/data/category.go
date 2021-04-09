package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/realotz/whole/api/cms/v1"
	"github.com/realotz/whole/internal/services/cms/biz"
	"github.com/realotz/whole/internal/services/cms/data/ent"
)

type entCategory ent.Category

// ent转换biz结构
func (p entCategory) BizStruct() *biz.Category {
	return &biz.Category{
		Id:         p.ID,
		Name:       p.Name,
		Pid:        p.Pid,
		Icon:       p.Icon,
		Desc:       p.Desc,
		UpdateTime: p.UpdateTime,
		CreateTime: p.CreateTime,
	}
}

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{data: data, log: log.NewHelper("category_repo", logger)}
}

//查询
func (ar *categoryRepo) GetCategory(ctx context.Context, id int64) (*biz.Category, error) {
	p, err := ar.data.db.Category.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entCategory(*p).BizStruct(), nil
}

//新增
func (ar *categoryRepo) CreateCategory(ctx context.Context, m *biz.Category) (*biz.Category, error) {
	modCreate := ar.data.db.Category.Create()
	p, err := modCreate.
		SetName(m.Name).
		SetPid(m.Pid).
		SetIcon(m.Icon).
		SetDesc(m.Desc).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return entCategory(*p).BizStruct(), err
}

//更新
func (ar *categoryRepo) UpdateCategory(ctx context.Context, id int64, m *biz.Category) (*biz.Category, error) {
	p, err := ar.data.db.Category.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	modUp := ar.data.db.Category.Update()
	if m.Name != "" && m.Name != p.Name {
		modUp.SetName(m.Name)
	}
	if m.Icon != "" && m.Icon != p.Icon {
		modUp.SetIcon(m.Icon)
	}
	if m.Desc != "" && m.Desc != p.Desc {
		modUp.SetDesc(m.Desc)
	}
	_, err = modUp.Save(ctx)
	return entCategory(*p).BizStruct(), err
}

//删除
func (ar *categoryRepo) DeleteCategory(ctx context.Context, ids []int64) error {
	tx, err := ar.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	for _, id := range ids {
		err = tx.Category.DeleteOneID(id).Exec(ctx)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// 列表搜索
func (ar *categoryRepo) ListCategory(ctx context.Context, op *pb.ListCategoryRequest) ([]*biz.Category, error) {
	query := ar.data.db.Category.Query()
	// todo 搜索条件
	ps, err := query.Order(ent.Desc("id")).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Category, 0)
	for _, p := range ps {
		rv = append(rv, entCategory(*p).BizStruct())
	}
	return rv, nil
}
