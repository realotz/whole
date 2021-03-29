package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/whole/internal/apps/users/biz"
	"github.com/realotz/whole/internal/apps/users/data/ent"
	"github.com/realotz/whole/internal/apps/users/data/ent/member"
)

type entMember ent.Member

func (p entMember) BizStudent() *biz.Member {
	return &biz.Member{
		ID:         p.ID,
		Sex:        p.Sex.String(),
		UUID:       p.UUID,
		Name:       p.Name,
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

type memberRepo struct {
	data *Data
	log  *log.Helper
}

// NewMemberRepo .
func NewMemberRepo(data *Data, logger log.Logger) biz.MemberRepo {
	return &memberRepo{
		data: data,
		log:  log.NewHelper("member_repo", logger),
	}
}

// 根据账户查询用户信息
// account 用户名 手机号 邮箱
func (ar *memberRepo) GetMemberForAccounts(ctx context.Context, account string) ([]*biz.Member, error) {
	ps, err := ar.data.db.Member.Query().Where(
		member.Or(
			member.Email(account),
			member.Mobile(account),
			member.Account(account),
		)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Member, 0)
	for _, p := range ps {
		rv = append(rv, entMember(*p).BizStudent())
	}
	return rv, nil
}

func (ar *memberRepo) ListMember(ctx context.Context) ([]*biz.Member, error) {
	ps, err := ar.data.db.Member.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Member, 0)
	for _, p := range ps {
		rv = append(rv, entMember(*p).BizStudent())
	}
	return rv, nil
}

func (ar *memberRepo) GetMember(ctx context.Context, id int64) (*biz.Member, error) {
	p, err := ar.data.db.Member.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entMember(*p).BizStudent(), nil
}

// 创建用户
func (ar *memberRepo) CreateMember(ctx context.Context, m *biz.Member) (*biz.Member, error) {
	p, err := ar.data.db.Member.Query().Where(member.Account(m.Account)).First(ctx)
	if err == nil && p != nil {
		return nil, fmt.Errorf("账号已存在")
	}
	modCreate := ar.data.db.Member.Create()
	if m.Sex != "" && member.SexValidator(member.Sex(m.Sex)) == nil {
		modCreate.SetSex(member.Sex(m.Sex))
	} else {
		modCreate.SetSex(member.SexUnknown)
	}
	if !m.Birthday.IsZero() {
		modCreate.SetBirthday(m.Birthday)
	}
	p, err = modCreate.
		SetAccount(m.Account).
		SetName(m.NickName).
		SetNickName(m.Name).
		SetEmail(m.Email).
		SetIDCard(m.IDCard).
		SetRole(m.Role).
		SetPassword(m.Password).
		SetSalt(m.Salt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return entMember(*p).BizStudent(), err
}

// 更新用户
func (ar *memberRepo) UpdateMember(ctx context.Context, id int64, m *biz.Member) (*biz.Member, error) {
	p, err := ar.data.db.Member.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	modUp := ar.data.db.Member.Update()
	if m.Sex != "" && member.SexValidator(member.Sex(m.Sex)) == nil {
		modUp.SetSex(member.Sex(m.Sex))
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
	return entMember(*p).BizStudent(), err
}

func (ar *memberRepo) DeleteMember(ctx context.Context, id int64) error {
	return ar.data.db.Member.DeleteOneID(id).Exec(ctx)
}
