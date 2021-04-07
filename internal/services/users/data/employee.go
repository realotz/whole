package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/whole/internal/services/users/biz"
	"github.com/realotz/whole/internal/services/users/data/ent"
	"github.com/realotz/whole/internal/services/users/data/ent/employee"
)

type entEmployee ent.Employee

// ent转换biz结构
func (p entEmployee) BizStruct() *biz.Employee {
	return &biz.Employee{
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

type employeeRepo struct {
	data *Data
	log  *log.Helper
}

// NewEmployeeRepo .
func NewEmployeeRepo(data *Data, logger log.Logger) biz.EmployeeRepo {
	return &employeeRepo{
		data: data,
		log:  log.NewHelper("employee_repo", logger),
	}
}

// 根据账户查询用户信息
// account 用户名 手机号 邮箱
func (ar *employeeRepo) GetEmployeeForAccounts(ctx context.Context, account string) ([]*biz.Employee, error) {
	ps, err := ar.data.db.Employee.Query().Where(
		employee.Or(
			employee.Email(account),
			employee.Mobile(account),
			employee.Account(account),
		)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Employee, 0)
	for _, p := range ps {
		rv = append(rv, entEmployee(*p).BizStruct())
	}
	return rv, nil
}

func (ar *employeeRepo) ListEmployee(ctx context.Context) ([]*biz.Employee, error) {
	ps, err := ar.data.db.Employee.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Employee, 0)
	for _, p := range ps {
		rv = append(rv, entEmployee(*p).BizStruct())
	}
	return rv, nil
}

func (ar *employeeRepo) GetEmployee(ctx context.Context, id int64) (*biz.Employee, error) {
	p, err := ar.data.db.Employee.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entEmployee(*p).BizStruct(), nil
}

// 创建用户
func (ar *employeeRepo) CreateEmployee(ctx context.Context, m *biz.Employee) (*biz.Employee, error) {
	p, err := ar.data.db.Employee.Query().Where(employee.Account(m.Account)).First(ctx)
	if err == nil && p != nil {
		return nil, fmt.Errorf("账号已存在")
	}
	modCreate := ar.data.db.Employee.Create()
	if m.Sex != "" && employee.SexValidator(employee.Sex(m.Sex)) == nil {
		modCreate.SetSex(employee.Sex(m.Sex))
	} else {
		modCreate.SetSex(employee.SexUnknown)
	}
	if !m.Birthday.IsZero() {
		modCreate.SetBirthday(m.Birthday)
	}
	p, err = modCreate.
		SetAccount(m.Account).
		SetName(m.NickName).
		SetNickName(m.Name).
		SetAvatar(m.Avatar).
		SetEmail(m.Email).
		SetIDCard(m.IDCard).
		SetRole(m.Role).
		SetPassword(m.Password).
		SetSalt(m.Salt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return entEmployee(*p).BizStruct(), err
}

// 更新用户
func (ar *employeeRepo) UpdateEmployee(ctx context.Context, id int64, m *biz.Employee) (*biz.Employee, error) {
	p, err := ar.data.db.Employee.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	modUp := ar.data.db.Employee.Update()
	if m.Sex != "" && employee.SexValidator(employee.Sex(m.Sex)) == nil {
		modUp.SetSex(employee.Sex(m.Sex))
	}
	if m.Name != "" && m.Name != p.Name {
		modUp.SetName(m.Name)
	}
	if m.Role != "" && m.Role != p.Role {
		modUp.SetRole(m.Role)
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
	return entEmployee(*p).BizStruct(), err
}

func (ar *employeeRepo) DeleteEmployee(ctx context.Context, id int64) error {
	return ar.data.db.Employee.DeleteOneID(id).Exec(ctx)
}
