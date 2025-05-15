package server

import (
	"context"
	"fmt"
	"github.com/bakhytzhanjzz/go-leetcode-platform/user-service/database"
	"github.com/bakhytzhanjzz/go-leetcode-platform/user-service/models"
	"github.com/bakhytzhanjzz/go-leetcode-platform/user-service/proto"

	"github.com/golang-jwt/jwt/v5"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	JWTSecret []byte
}

func (s *UserServer) GetUserByID(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	var user models.User
	if err := database.DB.First(&user, req.Id).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &proto.UserResponse{
		Id:       uint64(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *UserServer) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	token, _ := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return s.JWTSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint64(claims["user_id"].(float64))
		return &proto.ValidateTokenResponse{
			Valid:  true,
			UserId: userID,
		}, nil
	}

	return &proto.ValidateTokenResponse{
		Valid: false,
		Error: "Invalid token",
	}, nil
}
