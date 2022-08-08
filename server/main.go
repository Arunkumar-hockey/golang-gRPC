package main

import (
	"fmt"
	"golang-gRPC/server/controller"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Server started")

	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	controller.RegisterUserServiceServer(s, controller.NewUserControllerServer())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
