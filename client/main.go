package main

import (
	"context"
	"fmt"
	"log"

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
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to ChatService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)

	chat := chat_service.NewChatServiceClient(conn)
	if _, err := chat.GetAllUsers(context.Background(), &chat_service.GetAllUsersRequest{}); err != nil {
		return fmt.Errorf("failed to chat: %w", err)
	}

	log.Println("Successfully get list of users")
	return nil
}
