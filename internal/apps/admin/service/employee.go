package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	pb "github.com/realotz/whole/api/admin/v1"
	commPb "github.com/realotz/whole/api/comm"
	"github.com/realotz/whole/api/reason"
	"github.com/realotz/whole/internal/apps/admin/biz"
	"github.com/realotz/whole/pkg/token"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EmployeeService struct {
	pb.UnimplementedEmployeeServiceServer
	member *biz.EmployeeUsecase
}

func NewEmployeeService(member *biz.EmployeeUsecase) pb.EmployeeServiceServer {
	return &EmployeeService{
		member: member,
	}
}

//验证码登录
func (s *EmployeeService) LoginForCode(ctx context.Context, login *pb.EmployeeLogin) (*pb.EmployeeLoginRes, error) {
	panic("implement me")
}

//登出
func (s *EmployeeService) Logout(ctx context.Context, req *commPb.NullReq) (*commPb.NullReply, error) {
	panic("implement me")
}

//发送验证码
func (s *EmployeeService) Captcha(ctx context.Context, req *pb.CaptchaReq) (*commPb.NullReply, error) {
	panic("implement me")
}

// 获取当前登录用户id
func (s *EmployeeService) UserInfo(ctx context.Context, req *commPb.NullReq) (*pb.Employee, error) {
	userInfo := token.FormLoginContext(ctx)
	if userInfo == nil {
		return nil, errors.Unauthorized(reason.NotLogin, "未登录")
	}
	m, err := s.member.Get(ctx, userInfo.UserID)
	if err != nil {
		return nil, err
	}
	return s.protoMsg(m), nil
}

// 登录
func (s *EmployeeService) Login(ctx context.Context, req *pb.EmployeeLogin) (*pb.EmployeeLoginRes, error) {
	tk, etime, mb, err := s.member.Login(ctx, req.Account, req.Password)
	if err != nil {
		return nil, errors.Unauthorized(reason.LoginError, err.Error())
	}
	return &pb.EmployeeLoginRes{
		Token:          tk,
		ExpirationTime: etime,
		Employee:       s.protoMsg(mb),
	}, nil
}

//用户列表
func (s *EmployeeService) List(ctx context.Context, req *pb.EmployeeListOption) (*pb.EmployeeList, error) {
	ms, err := s.member.List(ctx)
	if err != nil {
		return nil, err
	}
	var resp = &pb.EmployeeList{List: make([]*pb.Employee, 0, len(ms))}
	for _, v := range ms {
		resp.List = append(resp.List, s.protoMsg(v))
	}
	return resp, nil
}

// 查询用户
func (s *EmployeeService) Get(ctx context.Context, req *pb.EmployeeGetOption) (*pb.Employee, error) {
	m, err := s.member.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return s.protoMsg(m), nil
}

// 创建用户
func (s *EmployeeService) Create(ctx context.Context, req *pb.EmployeeOption) (*pb.Employee, error) {
	m, err := s.member.Create(ctx, &biz.Employee{
		Sex:      req.Sex,
		Name:     req.Name,
		Account:  req.Account,
		Avatar:   req.Avatar,
		Email:    req.Email,
		Mobile:   req.Mobile,
		IDCard:   req.IdCard,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return s.protoMsg(m), nil
}

// 更新用户
func (s *EmployeeService) Update(ctx context.Context, req *pb.EmployeeOption) (*pb.Employee, error) {
	m, err := s.member.Update(ctx, req.Id, &biz.Employee{
		Sex:      req.Sex,
		Name:     req.Name,
		Account:  req.Account,
		Email:    req.Email,
		Mobile:   req.Mobile,
		IDCard:   req.IdCard,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return s.protoMsg(m), nil
}

// 删除用户
func (s *EmployeeService) Delete(ctx context.Context, req *pb.EmployeeDeleteOption) (*commPb.NullReply, error) {
	err := s.member.Delete(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	return &commPb.NullReply{}, nil
}

// biz转换pb
func (uc *EmployeeService) protoMsg(m *biz.Employee) *pb.Employee {
	msg := &pb.Employee{
		Id:         m.ID,
		Uuid:       m.UUID.String(),
		Name:       m.Name,
		Account:    m.Account,
		Email:      m.Email,
		Mobile:     m.Mobile,
		IdCard:     m.IDCard,
		Sex:        m.Sex,
		LastIp:     m.LastIP,
		CreateTime: timestamppb.New(m.CreateTime),
		UpdateTime: timestamppb.New(m.UpdateTime),
	}
	if !m.Birthday.IsZero() {
		msg.Birthday = timestamppb.New(m.Birthday)
	}
	if !m.LastTime.IsZero() {
		msg.LastTime = timestamppb.New(m.LastTime)
	}
	return msg
}
