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
	pb "github.com/realotz/whole/api/admin/v1"
	"github.com/realotz/whole/pkg/utils"
	"time"
)

// 用户biz结构
type Employee struct {
	Sex     string `json:"sex,omitempty"`
	Name    string `json:"name,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	Account string `json:"account,omitempty"`
	Email   string `json:"email,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
	IDCard  string `json:"id_card,omitempty"`

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

type EmployeeRepo interface {
	ListEmployee(ctx context.Context, opt *pb.EmployeeListOption) ([]*Employee, int64, error)
	GetEmployee(ctx context.Context, id int64) (*Employee, error)
	CreateEmployee(ctx context.Context, article *Employee) (*Employee, error)
	UpdateEmployee(ctx context.Context, id int64, article *Employee) (*Employee, error)
	DeleteEmployee(ctx context.Context, id []int64) error
	GetEmployeeForAccounts(ctx context.Context, account string) ([]*Employee, error)
	// 创建jwttoken
	CreateToken(member *Employee, time int64) (string, error)
}

type EmployeeUsecase struct {
	repo EmployeeRepo
	log  *log.Helper
}

func NewEmployeeUsecase(repo EmployeeRepo, logger log.Logger) *EmployeeUsecase {
	return &EmployeeUsecase{repo: repo, log: log.NewHelper("biz/member", logger)}
}

// 查询列表
func (uc *EmployeeUsecase) List(ctx context.Context, opt *pb.EmployeeListOption) (ps []*Employee, total int64, err error) {
	return uc.repo.ListEmployee(ctx, opt)
}

// 根据主键查询
func (uc *EmployeeUsecase) Get(ctx context.Context, id int64) (p *Employee, err error) {
	return uc.repo.GetEmployee(ctx, id)
}

// 创建用户
func (uc *EmployeeUsecase) Create(ctx context.Context, member *Employee) (*Employee, error) {
	if member.Account == "" {
		return nil, fmt.Errorf("账号不能为空")
	}
	if member.Password == "" {
		member.Password = "Admin!@#321"
	}
	member.SetPassword(member.Password)
	mem, err := uc.repo.CreateEmployee(ctx, member)
	if err != nil {
		return nil, err
	}
	// todo 发送注册邮件
	return mem, err
}

// 更新用户
func (uc *EmployeeUsecase) Update(ctx context.Context, id int64, member *Employee) (mem *Employee, err error) {
	if member.Password != "" {
		member.SetPassword(member.Password)
		defer func() {
			if err == nil {
				//todo 发送密码变动邮件
			}
		}()
	}
	return uc.repo.UpdateEmployee(ctx, id, member)
}

// 删除用户
func (uc *EmployeeUsecase) Delete(ctx context.Context, ids []int64) error {
	return uc.repo.DeleteEmployee(ctx, ids)
}

// 登录
func (uc *EmployeeUsecase) Login(ctx context.Context, account, password string) (token string, expireTime int64, m *Employee, err error) {
	members, err := uc.repo.GetEmployeeForAccounts(ctx, account)
	if err != nil {
		return "", 0, nil, err
	}
	for _, mb := range members {
		if mb.CheckPassword(password) {
			_, _ = uc.repo.UpdateEmployee(ctx, mb.ID, &Employee{
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
func (uc *Employee) CheckPassword(password string) bool {
	return uc.GenPassword(password) == uc.Password
}

// 创建密码
func (uc *Employee) SetPassword(password string) {
	uc.Salt = utils.RandomString(32)
	uc.Password = uc.GenPassword(password)
}

// 创建密码
func (uc *Employee) GenPassword(password string) string {
	hmacEnt := hmac.New(md5.New, []byte(uc.Salt))
	hmacEnt.Write([]byte(password))
	return hex.EncodeToString(hmacEnt.Sum([]byte("")))
}
