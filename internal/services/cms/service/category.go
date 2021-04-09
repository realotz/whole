package service

import (
	"context"
	pb "github.com/realotz/whole/api/cms/v1"
	"github.com/realotz/whole/internal/services/cms/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewCategoryService(bz *biz.CategoryUsecase) pb.CategoryServiceServer {
	return &CategoryService{
		category: bz,
	}
}

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	category *biz.CategoryUsecase
}

//创建文章类目
func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.Category) (*pb.CreateCategoryReply, error) {
	m, err := s.category.Create(ctx, &biz.Category{
		Name: req.Name,
		Pid:  req.Pid,
		Icon: req.Icon,
		Desc: req.Desc,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateCategoryReply{Data: s.protoMsg(m)}, nil
}

//更新文章类目
func (s *CategoryService) UpdateCategory(ctx context.Context, req *pb.Category) (*pb.UpdateCategoryReply, error) {
	m, err := s.category.Update(ctx, &biz.Category{
		Id:   req.Id,
		Name: req.Name,
		Pid:  req.Pid,
		Icon: req.Icon,
		Desc: req.Desc,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCategoryReply{Data: s.protoMsg(m)}, nil
}

//批量删除文章类目
func (s *CategoryService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryReply, error) {
	err := s.category.Delete(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCategoryReply{}, nil
}

//获取文章类目详情
func (s *CategoryService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	m, err := s.category.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return s.protoMsg(m), nil
}

//查询文章类目列表
func (s *CategoryService) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryReply, error) {
	ms, err := s.category.List(ctx, req)
	if err != nil {
		return nil, err
	}
	var resp = &pb.ListCategoryReply{List: make([]*pb.Category, 0, len(ms))}
	for _, v := range ms {
		resp.List = append(resp.List, s.protoMsg(v))
	}
	return resp, nil
}

//biz转换pb
func (uc *CategoryService) protoMsg(m *biz.Category) *pb.Category {
	msg := &pb.Category{
		Id:   m.Id,
		Name: m.Name,
		Pid:  m.Pid,
		Icon: m.Icon,
		Desc: m.Desc,
	}
	if !m.CreateTime.IsZero() {
		msg.CreateTime = timestamppb.New(m.CreateTime)
	}
	if !m.UpdateTime.IsZero() {
		msg.UpdateTime = timestamppb.New(m.UpdateTime)
	}
	return msg
}
