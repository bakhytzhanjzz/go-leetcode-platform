package grpcclient

import (
	"context"
	"log"
	"time"

	problempb "github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/proto"
	"google.golang.org/grpc"
)

type ProblemClient struct {
	client problempb.ProblemServiceClient
}

func NewProblemClient(addr string) *ProblemClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to ProblemService at %s: %v", addr, err)
	}

	client := problempb.NewProblemServiceClient(conn)
	return &ProblemClient{client: client}
}

func (pc *ProblemClient) GetProblem(id uint64) (*problempb.ProblemResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return pc.client.GetProblemByID(ctx, &problempb.GetProblemRequest{Id: id})
}
