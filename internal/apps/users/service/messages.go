package service

import (
	"context"

	pb "github.com/realotz/whole/api/users/v1"
)

type MessageService struct {
	pb.UnimplementedMessageServiceServer
}

func NewMessageService() pb.MessageServiceServer {
	return &MessageService{}
}

func (s *MessageService) List(ctx context.Context, req *pb.MessageListOption) (*pb.MessageList, error) {
	return &pb.MessageList{}, nil
}
func (s *MessageService) Get(ctx context.Context, req *pb.MessageGetOption) (*pb.Message, error) {
	return &pb.Message{}, nil
}
func (s *MessageService) Create(ctx context.Context, req *pb.MessageGetOption) (*pb.Message, error) {
	return &pb.Message{}, nil
}
func (s *MessageService) Update(ctx context.Context, req *pb.MessageUpdateOption) (*pb.Message, error) {
	return &pb.Message{}, nil
}
func (s *MessageService) Delete(ctx context.Context, req *pb.MessageDeleteOption) (*pb.Message, error) {
	return &pb.Message{}, nil
}
