package grpcclient

import (
	"context"
	"log"
	"time"

	userpb "github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/proto"
	"google.golang.org/grpc"
)

type UserClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(addr string) *UserClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to UserService at %s: %v", addr, err)
	}

	client := userpb.NewUserServiceClient(conn)
	return &UserClient{client: client}
}

func (uc *UserClient) ValidateToken(token string) (uint64, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := uc.client.ValidateToken(ctx, &userpb.ValidateTokenRequest{Token: token})
	if err != nil {
		return 0, false, err.Error()
	}

	return res.UserId, res.Valid, res.Error
}
