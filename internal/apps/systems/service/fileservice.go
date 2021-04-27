package service

import (
	"context"
	"github.com/realotz/whole/internal/apps/systems/biz"

	pb "github.com/realotz/whole/api/systems/v1"
)

type FileServiceService struct {
	pb.UnimplementedFileServiceServer
	file *biz.FileUsecase
}

func NewFileServiceService(file *biz.FileUsecase) pb.FileServiceServer {
	return &FileServiceService{file: file}
}

func (s *FileServiceService) List(ctx context.Context, req *pb.FileListOption) (*pb.FileList, error) {
	return &pb.FileList{}, nil
}
func (s *FileServiceService) Get(ctx context.Context, req *pb.FileGetOption) (*pb.File, error) {
	return &pb.File{}, nil
}
func (s *FileServiceService) Create(ctx context.Context, req *pb.FileCreateOption) (*pb.File, error) {
	return &pb.File{}, nil
}
func (s *FileServiceService) Update(ctx context.Context, req *pb.FileUpdateOption) (*pb.File, error) {
	return &pb.File{}, nil
}
func (s *FileServiceService) Delete(ctx context.Context, req *pb.FileDeleteOption) (*pb.File, error) {
	return &pb.File{}, nil
}
