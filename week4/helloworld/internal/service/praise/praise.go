package service

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    pb "helloworld/api/praise"
    "helloworld/internal/biz"
)

type PraiseService struct {
    pb.UnimplementedPraiseServer
    uc  *biz.GreeterUsecase
    log *log.Helper
}

func NewPraiseService() *PraiseService {
    return &PraiseService{}
}

func (s *PraiseService) AddPraise(ctx context.Context, req *pb.AddPraiseRequest) (*pb.AddPraiseReply, error) {
    return &pb.AddPraiseReply{}, nil
}
func (s *PraiseService) CancelPraise(ctx context.Context, req *pb.CancelPraiseRequest) (*pb.CancelPraiseReply, error) {
    return &pb.CancelPraiseReply{}, nil
}
func (s *PraiseService) IsPraisePraise(ctx context.Context, req *pb.IsPraisePraiseRequest) (*pb.IsPraisePraiseReply, error) {
    return &pb.IsPraisePraiseReply{}, nil
}
