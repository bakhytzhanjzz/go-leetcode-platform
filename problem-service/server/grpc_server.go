package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ProblemServer struct {
	proto.UnimplementedProblemServiceServer
	DB *gorm.DB
}

func (s *ProblemServer) GetProblemByID(ctx context.Context, req *proto.GetProblemRequest) (*proto.ProblemResponse, error) {
	var problem models.Problem
	if err := s.DB.First(&problem, req.Id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("problem not found")
		}
		return nil, fmt.Errorf("internal error: %v", err)
	}

	return &proto.ProblemResponse{
		Id:          uint64(problem.ID),
		Title:       problem.Title,
		Description: problem.Description,
		Difficulty:  problem.Difficulty,
	}, nil
}

func StartGRPCServer(db *gorm.DB) {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterProblemServiceServer(grpcServer, &ProblemServer{DB: db})

	log.Println("ProblemService gRPC server listening on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
