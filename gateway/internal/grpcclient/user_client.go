package grpcclient

import (
	"context"
	"log"

	userpb "github.com/bakhytzhanjzz/go-leetcode-platform/user-service/proto"
	"google.golang.org/grpc"
)

type UserClient struct {
	Client userpb.UserServiceClient
}

func NewUserClient(address string) *UserClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	client := userpb.NewUserServiceClient(conn)
	return &UserClient{Client: client}
}

func (uc *UserClient) ValidateToken(token string) (*userpb.ValidateTokenResponse, error) {
	return uc.Client.ValidateToken(context.Background(), &userpb.ValidateTokenRequest{
		Token: token,
	})
}
