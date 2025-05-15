package main

import (
	"github.com/bakhytzhanjzz/go-leetcode-platform/user-service/database"
	"github.com/bakhytzhanjzz/go-leetcode-platform/user-service/handlers"
	"github.com/bakhytzhanjzz/go-leetcode-platform/user-service/proto"
	"github.com/bakhytzhanjzz/go-leetcode-platform/user-service/server"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	database.Init()

	// Start gRPC server in a goroutine
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		userService := &server.UserServer{JWTSecret: []byte("your_secret_key")}
		proto.RegisterUserServiceServer(grpcServer, userService)
		log.Println("User gRPC server listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Start REST API
	r := gin.Default()
	r.POST("/users/register", handlers.Register)
	r.POST("/users/login", handlers.Login)
	r.GET("/users/:id", handlers.GetUser)
	r.Run(":8083")
}
