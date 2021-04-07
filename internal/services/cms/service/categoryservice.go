package service

import (
	"context"
	"github.com/realotz/whole/internal/services/cms/biz"

	pb "github.com/realotz/whole/api/cms/v1"
)

type CategoryServiceService struct {
	pb.UnimplementedCategoryServiceServer
	category *biz.CategoryUsecase
}

func NewCategoryServiceService(category *biz.CategoryUsecase) pb.CategoryServiceServer {
	return &CategoryServiceService{
		category: category,
	}
}

func (s *CategoryServiceService) List(ctx context.Context, req *pb.CategoryListOption) (*pb.CategoryList, error) {
	return &pb.CategoryList{}, nil
}
func (s *CategoryServiceService) Get(ctx context.Context, req *pb.CategoryGetOption) (*pb.Category, error) {
	return &pb.Category{}, nil
}
func (s *CategoryServiceService) Create(ctx context.Context, req *pb.CategoryGetOption) (*pb.Category, error) {
	return &pb.Category{}, nil
}
func (s *CategoryServiceService) Update(ctx context.Context, req *pb.CategoryUpdateOption) (*pb.Category, error) {
	return &pb.Category{}, nil
}
func (s *CategoryServiceService) Delete(ctx context.Context, req *pb.CategoryDeleteOption) (*pb.Category, error) {
	return &pb.Category{}, nil
}
