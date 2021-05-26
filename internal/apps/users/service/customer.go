package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	coomPb "github.com/realotz/whole/api/comm"
	"github.com/realotz/whole/api/reason"
	pb "github.com/realotz/whole/api/users/v1"
	"github.com/realotz/whole/internal/apps/users/biz"
	"github.com/realotz/whole/pkg/token"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewCustomerService(member *biz.CustomerUsecase) pb.CustomerServiceServer {
	return &CustomerService{
		member: member,
		domain: "Customer",
	}
}

type CustomerService struct {
	pb.UnimplementedCustomerServiceServer
	member *biz.CustomerUsecase
	domain string
}

func (c CustomerService) LoginForCode(ctx context.Context, login *pb.CustomerLogin) (*pb.CustomerLoginRes, error) {
	panic("implement me")
}

func (c CustomerService) Logout(ctx context.Context, req *coomPb.NullReq) (*coomPb.NullReply, error) {
	panic("implement me")
}

func (c CustomerService) Captcha(ctx context.Context, req *pb.CaptchaReq) (*coomPb.NullReply, error) {
	panic("implement me")
}

// 获取当前登录用户id
func (s *CustomerService) UserInfo(ctx context.Context, req *coomPb.NullReq) (*pb.Customer, error) {
	userInfo := token.FormLoginContext(ctx)
	if userInfo == nil {
		return nil, errors.Unauthorized(s.domain, reason.NotLogin, "未登录")
	}
	m, err := s.member.Get(ctx, userInfo.UserID)
	if err != nil {
		return nil, err
	}
	return s.protoMsg(m), nil
}

// 登录
func (s *CustomerService) Login(ctx context.Context, req *pb.CustomerLogin) (*pb.CustomerLoginRes, error) {
	tk, etime, mb, err := s.member.Login(ctx, req.Account, req.Password)
	if err != nil {
		return nil, errors.Unauthorized(s.domain, reason.LoginError, err.Error())
	}
	return &pb.CustomerLoginRes{
		Token:          tk,
		ExpirationTime: etime,
		Customer:       s.protoMsg(mb),
	}, nil
}

//用户列表
func (s *CustomerService) List(ctx context.Context, req *pb.CustomerListOption) (*pb.CustomerList, error) {
	ms, err := s.member.List(ctx, req)
	if err != nil {
		return nil, err
	}
	var resp = &pb.CustomerList{List: make([]*pb.Customer, 0, len(ms))}
	for _, v := range ms {
		resp.List = append(resp.List, s.protoMsg(v))
	}
	return resp, nil
}

// 查询用户
func (s *CustomerService) Get(ctx context.Context, req *pb.CustomerGetOption) (*pb.Customer, error) {
	m, err := s.member.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return s.protoMsg(m), nil
}

// 创建用户
func (s *CustomerService) Create(ctx context.Context, req *pb.CustomerOption) (*pb.Customer, error) {
	m, err := s.member.Create(ctx, &biz.Customer{
		Sex:      req.Sex,
		Name:     req.Name,
		Account:  req.Account,
		Avatar:   req.Avatar,
		NickName: req.NickName,
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
func (s *CustomerService) Update(ctx context.Context, req *pb.CustomerOption) (*pb.Customer, error) {
	m, err := s.member.Update(ctx, req.Id, &biz.Customer{
		Sex:      req.Sex,
		Name:     req.Name,
		Account:  req.Account,
		NickName: req.NickName,
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
func (s *CustomerService) Delete(ctx context.Context, req *pb.CustomerDeleteOption) (*coomPb.NullReply, error) {
	err := s.member.Delete(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	return &coomPb.NullReply{}, nil
}

// biz转换pb
func (uc *CustomerService) protoMsg(m *biz.Customer) *pb.Customer {
	msg := &pb.Customer{
		Id:         m.ID,
		Uuid:       m.UUID.String(),
		Name:       m.Name,
		Account:    m.Account,
		NickName:   m.NickName,
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
