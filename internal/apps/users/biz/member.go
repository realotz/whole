package biz

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/realotz/whole/pkg/utils"
	"time"
)

type Member struct {
	Sex      string `json:"sex,omitempty"`
	Name     string `json:"name,omitempty"`
	Role     string `json:"role,omitempty"`
	Account  string `json:"account,omitempty"`
	NickName string `json:"nick_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	IDCard   string `json:"id_card,omitempty"`

	ID         int64     `json:"id,omitempty"`
	Password   string    `json:"-"`
	Salt       string    `json:"-"`
	LastIP     string    `json:"last_ip,omitempty"`
	Birthday   time.Time `json:"birthday,omitempty"`
	LastTime   time.Time `json:"last_time,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
	UUID       uuid.UUID `json:"uuid,omitempty"`
}

type MemberRepo interface {
	ListMember(ctx context.Context) ([]*Member, error)
	GetMember(ctx context.Context, id int64) (*Member, error)
	CreateMember(ctx context.Context, article *Member) (*Member, error)
	UpdateMember(ctx context.Context, id int64, article *Member) (*Member, error)
	DeleteMember(ctx context.Context, id int64) error
	GetMemberForAccounts(ctx context.Context, account string) ([]*Member, error)
	// 创建jwttoken
	CreateToken(member *Member, time int64) (string, error)
}

type MemberUsecase struct {
	repo MemberRepo
	log  *log.Helper
}

func NewMemberUsecase(repo MemberRepo, logger log.Logger) *MemberUsecase {
	return &MemberUsecase{repo: repo, log: log.NewHelper("biz/member", logger)}
}

func (uc *MemberUsecase) List(ctx context.Context) (ps []*Member, err error) {
	return uc.repo.ListMember(ctx)
}

func (uc *MemberUsecase) Get(ctx context.Context, id int64) (p *Member, err error) {
	return uc.repo.GetMember(ctx, id)
}

// 创建用户
func (uc *MemberUsecase) Create(ctx context.Context, member *Member) (*Member, error) {
	if member.Account == "" {
		return nil, fmt.Errorf("账号不能为空")
	}
	if member.Password == "" {
		return nil, fmt.Errorf("密码不能为空")
	}
	member.SetPassword(member.Password)
	return uc.repo.CreateMember(ctx, member)
}

// 更新用户
func (uc *MemberUsecase) Update(ctx context.Context, id int64, member *Member) (*Member, error) {
	if member.Password != "" {
		member.SetPassword(member.Password)
	}
	return uc.repo.UpdateMember(ctx, id, member)
}

// 删除用户
func (uc *MemberUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteMember(ctx, id)
}

// 登录
func (uc *MemberUsecase) Login(ctx context.Context, account, password string) (token string, expireTime int64, m *Member, err error) {
	members, err := uc.repo.GetMemberForAccounts(ctx, account)
	if err != nil {
		return "", 0, nil, err
	}
	for _, mb := range members {
		if mb.CheckPassword(password) {
			_, _ = uc.repo.UpdateMember(ctx, mb.ID, &Member{
				LastIP:   utils.GetClientIp(ctx),
				Birthday: time.Now(),
			})
			expireTime = time.Now().Add(time.Hour * 24 * 3).Unix()
			token, err = uc.repo.CreateToken(mb, expireTime)
			return token, expireTime, mb, err
		}
	}
	return "", 0, nil, errors.New("账号或密码错误")
}

// 检查密码
func (uc *Member) CheckPassword(password string) bool {
	return uc.GenPassword(password) == uc.Password
}

// 创建密码
func (uc *Member) SetPassword(password string) {
	uc.Salt = utils.RandomString(32)
	uc.Password = uc.GenPassword(password)
}

// 创建密码
func (uc *Member) GenPassword(password string) string {
	hmacEnt := hmac.New(md5.New, []byte(uc.Salt))
	hmacEnt.Write([]byte(password))
	return hex.EncodeToString(hmacEnt.Sum([]byte("")))
}
