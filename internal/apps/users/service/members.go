package service

import (
	"context"
	pb "github.com/realotz/whole/api/users/v1"
	"github.com/realotz/whole/internal/apps/users/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MemberService struct {
	pb.UnimplementedMemberServiceServer
	member *biz.MemberUsecase
}

func NewMemberService(member *biz.MemberUsecase) pb.MemberServiceServer {
	return &MemberService{
		member: member,
	}
}

//用户列表
func (s *MemberService) List(ctx context.Context, req *pb.MemberListOption) (*pb.MemberList, error) {
	ms, err := s.member.List(ctx)
	if err != nil {
		return nil, err
	}
	var resp = &pb.MemberList{List: make([]*pb.Member, 0, len(ms))}
	for _, v := range ms {
		resp.List = append(resp.List, s.protoMsg(v))
	}
	return resp, nil
}

// 查询用户
func (s *MemberService) Get(ctx context.Context, req *pb.MemberGetOption) (*pb.Member, error) {
	m, err := s.member.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return s.protoMsg(m), nil
}

// 创建用户
func (s *MemberService) Create(ctx context.Context, req *pb.MemberCreateOption) (*pb.Member, error) {
	m, err := s.member.Create(ctx, &biz.Member{
		Sex:      req.Sex,
		Name:     req.Name,
		Role:     req.Role,
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

// 更新用户
func (s *MemberService) Update(ctx context.Context, req *pb.MemberUpdateOption) (*pb.Member, error) {
	m, err := s.member.Update(ctx, req.Id, &biz.Member{
		Sex:      req.Sex,
		Name:     req.Name,
		Role:     req.Role,
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
func (s *MemberService) Delete(ctx context.Context, req *pb.MemberDeleteOption) (*pb.NullReply, error) {
	err := s.member.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.NullReply{}, nil
}

// 登录
func (s *MemberService) Login(ctx context.Context, req *pb.MemberLogin) (*pb.MemberLoginRes, error) {
	tk, etime, mb, err := s.member.Login(ctx, req.Account, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.MemberLoginRes{
		Token:          tk,
		ExpirationTime: etime,
		Member:         s.protoMsg(mb),
	}, nil
}

// biz转换pb
func (uc *MemberService) protoMsg(m *biz.Member) *pb.Member {
	msg := &pb.Member{
		Id:         m.ID,
		Uuid:       m.UUID.String(),
		Name:       m.Name,
		Account:    m.Account,
		NickName:   m.NickName,
		Role:       m.Role,
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
