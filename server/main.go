package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	chat_service "github.com/gregb122/grpc-chat-protobufs/gen/go/chat_service/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	chat_service.RegisterChatServiceServer(server, &chatServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// chatServiceServer implements the ChatService API.
type chatServiceServer struct {
	chat_service.UnimplementedChatServiceServer
}

// Get all users
func (s *chatServiceServer) GetAllUsers(ctx context.Context, req *chat_service.GetAllUsersRequest) (*chat_service.GetAllUsersResponse, error) {
	log.Println("Got a request to get users")

	return &chat_service.GetAllUsersResponse{}, nil
}
