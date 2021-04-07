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

type Customer struct {
	Sex      string `json:"sex,omitempty"`
	Name     string `json:"name,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
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

type CustomerRepo interface {
	ListCustomer(ctx context.Context) ([]*Customer, error)
	GetCustomer(ctx context.Context, id int64) (*Customer, error)
	CreateCustomer(ctx context.Context, article *Customer) (*Customer, error)
	UpdateCustomer(ctx context.Context, id int64, article *Customer) (*Customer, error)
	DeleteCustomer(ctx context.Context, id int64) error
	GetCustomerForAccounts(ctx context.Context, account string) ([]*Customer, error)
	// 创建jwttoken
	CreateToken(member *Customer, time int64) (string, error)
}

type CustomerUsecase struct {
	repo CustomerRepo
	log  *log.Helper
}

func NewCustomerUsecase(repo CustomerRepo, logger log.Logger) *CustomerUsecase {
	return &CustomerUsecase{repo: repo, log: log.NewHelper("biz/member", logger)}
}

func (uc *CustomerUsecase) List(ctx context.Context) (ps []*Customer, err error) {
	return uc.repo.ListCustomer(ctx)
}

func (uc *CustomerUsecase) Get(ctx context.Context, id int64) (p *Customer, err error) {
	return uc.repo.GetCustomer(ctx, id)
}

// 创建用户
func (uc *CustomerUsecase) Create(ctx context.Context, member *Customer) (*Customer, error) {
	if member.Account == "" {
		return nil, fmt.Errorf("账号不能为空")
	}
	if member.Password == "" {
		return nil, fmt.Errorf("密码不能为空")
	}
	member.SetPassword(member.Password)
	return uc.repo.CreateCustomer(ctx, member)
}

// 更新用户
func (uc *CustomerUsecase) Update(ctx context.Context, id int64, member *Customer) (*Customer, error) {
	if member.Password != "" {
		member.SetPassword(member.Password)
	}
	return uc.repo.UpdateCustomer(ctx, id, member)
}

// 删除用户
func (uc *CustomerUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteCustomer(ctx, id)
}

// 登录
func (uc *CustomerUsecase) Login(ctx context.Context, account, password string) (token string, expireTime int64, m *Customer, err error) {
	members, err := uc.repo.GetCustomerForAccounts(ctx, account)
	if err != nil {
		return "", 0, nil, err
	}
	for _, mb := range members {
		if mb.CheckPassword(password) {
			_, _ = uc.repo.UpdateCustomer(ctx, mb.ID, &Customer{
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
func (uc *Customer) CheckPassword(password string) bool {
	return uc.GenPassword(password) == uc.Password
}

// 创建密码
func (uc *Customer) SetPassword(password string) {
	uc.Salt = utils.RandomString(32)
	uc.Password = uc.GenPassword(password)
}

// 创建密码
func (uc *Customer) GenPassword(password string) string {
	hmacEnt := hmac.New(md5.New, []byte(uc.Salt))
	hmacEnt.Write([]byte(password))
	return hex.EncodeToString(hmacEnt.Sum([]byte("")))
}
