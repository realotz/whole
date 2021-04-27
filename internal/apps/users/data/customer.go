package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/realotz/whole/api/users/v1"
	"github.com/realotz/whole/internal/apps/users/biz"
	"github.com/realotz/whole/internal/apps/users/data/ent"
	"github.com/realotz/whole/internal/apps/users/data/ent/customer"
)

type entCustomer ent.Customer

// ent转换biz结构
func (p entCustomer) BizStruct() *biz.Customer {
	return &biz.Customer{
		ID:         p.ID,
		Sex:        p.Sex.String(),
		UUID:       p.UUID,
		Name:       p.Name,
		Avatar:     p.Avatar,
		Account:    p.Account,
		Birthday:   p.Birthday,
		NickName:   p.NickName,
		Email:      p.Email,
		Mobile:     p.Mobile,
		IDCard:     p.IDCard,
		Password:   p.Password,
		Salt:       p.Salt,
		LastIP:     p.LastIP,
		LastTime:   p.LastTime,
		CreateTime: p.CreateTime,
		UpdateTime: p.UpdateTime,
	}
}

type customerRepo struct {
	data *Data
	log  *log.Helper
}

// NewCustomerRepo .
func NewCustomerRepo(data *Data, logger log.Logger) biz.CustomerRepo {
	return &customerRepo{
		data: data,
		log:  log.NewHelper("customer_repo", logger),
	}
}

// 根据账户查询用户信息
// account 用户名 手机号 邮箱
func (ar *customerRepo) GetCustomerForAccounts(ctx context.Context, account string) ([]*biz.Customer, error) {
	ps, err := ar.data.db.Customer.Query().Where(
		customer.Or(
			customer.Email(account),
			customer.Mobile(account),
			customer.Account(account),
		)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Customer, 0)
	for _, p := range ps {
		rv = append(rv, entCustomer(*p).BizStruct())
	}
	return rv, nil
}

// 用户列表搜索
func (ar *customerRepo) ListCustomer(ctx context.Context, op *pb.CustomerListOption) ([]*biz.Customer, error) {
	query := ar.data.db.Customer.Query()
	// 关键字搜索
	if op.Keyword != "" {
		query.Where(customer.Or(
			customer.NameContains("%"+op.Keyword+"%"),
			customer.EmailContains("%"+op.Keyword+"%"),
			customer.MobileContains("%"+op.Keyword+"%"),
			customer.IDCardContains("%"+op.Keyword+"%"),
			customer.NickNameContains("%"+op.Keyword+"%"),
		))
	}
	//性别查询
	if op.Sex != "" {
		if err := customer.SexValidator(customer.Sex(op.Sex)); err != nil {
			return nil, err
		}
		query.Where(customer.SexEQ(customer.Sex(op.Sex)))
	}
	// 电子邮件
	if op.Email != "" {
		query.Where(customer.EmailEQ(op.Email))
	}
	// 手机号
	if op.Mobile != "" {
		query.Where(customer.MobileEQ(op.Mobile))
	}
	// 身份证
	if op.IdCard != "" {
		query.Where(customer.IDCardEQ(op.IdCard))
	}
	// 名字
	if op.Name != "" {
		query.Where(customer.Or(
			customer.NameContains("%"+op.Name+"%"),
			customer.NickNameContains("%"+op.Name+"%"),
		))
	}
	ps, err := query.Order(ent.Desc("id")).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Customer, 0)
	for _, p := range ps {
		rv = append(rv, entCustomer(*p).BizStruct())
	}
	return rv, nil
}

//根据id获取用户
func (ar *customerRepo) GetCustomer(ctx context.Context, id int64) (*biz.Customer, error) {
	p, err := ar.data.db.Customer.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entCustomer(*p).BizStruct(), nil
}

// 创建用户
func (ar *customerRepo) CreateCustomer(ctx context.Context, m *biz.Customer) (*biz.Customer, error) {
	p, err := ar.data.db.Customer.Query().Where(customer.Account(m.Account)).First(ctx)
	if err == nil && p != nil {
		return nil, fmt.Errorf("账号已存在")
	}
	modCreate := ar.data.db.Customer.Create()
	if m.Sex != "" && customer.SexValidator(customer.Sex(m.Sex)) == nil {
		modCreate.SetSex(customer.Sex(m.Sex))
	} else {
		modCreate.SetSex(customer.SexUnknown)
	}
	if !m.Birthday.IsZero() {
		modCreate.SetBirthday(m.Birthday)
	}
	p, err = modCreate.
		SetAccount(m.Account).
		SetName(m.Name).
		SetNickName(m.NickName).
		SetAvatar(m.Avatar).
		SetEmail(m.Email).
		SetIDCard(m.IDCard).
		SetPassword(m.Password).
		SetSalt(m.Salt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return entCustomer(*p).BizStruct(), err
}

// 更新用户
func (ar *customerRepo) UpdateCustomer(ctx context.Context, id int64, m *biz.Customer) (*biz.Customer, error) {
	p, err := ar.data.db.Customer.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	modUp := ar.data.db.Customer.Update()
	if m.Sex != "" && customer.SexValidator(customer.Sex(m.Sex)) == nil {
		modUp.SetSex(customer.Sex(m.Sex))
	}
	if m.Name != "" && m.Name != p.Name {
		modUp.SetName(m.Name)
	}
	if m.Account != "" && m.Account != p.Account {
		modUp.SetAccount(m.Account)
	}
	if m.Avatar != "" && m.Avatar != p.Avatar {
		modUp.SetAvatar(m.Avatar)
	}
	if m.NickName != "" && m.NickName != p.NickName {
		modUp.SetNickName(m.NickName)
	}
	if !m.Birthday.IsZero() {
		modUp.SetBirthday(m.Birthday)
	}
	if m.Email != "" && m.Email != p.Email {
		modUp.SetEmail(m.Email)
	}
	if m.Mobile != "" && m.Mobile != p.Mobile {
		modUp.SetMobile(m.Mobile)
	}
	if m.IDCard != "" && m.IDCard != p.IDCard {
		modUp.SetIDCard(m.IDCard)
	}
	if m.Salt != "" && m.Salt != p.Salt {
		modUp.SetSalt(m.Salt)
	}
	if m.Password != "" && m.Password != p.Password {
		modUp.SetPassword(m.Password)
	}
	if m.LastIP != "" && m.LastIP != p.LastIP {
		modUp.SetIDCard(m.LastIP)
	}
	if !m.LastTime.IsZero() {
		modUp.SetLastTime(m.LastTime)
	}
	_, err = modUp.Save(ctx)
	return entCustomer(*p).BizStruct(), err
}

func (ar *customerRepo) DeleteCustomer(ctx context.Context, ids []int64) error {
	tx, err := ar.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	for _, id := range ids {
		err = tx.Customer.DeleteOneID(id).Exec(ctx)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
