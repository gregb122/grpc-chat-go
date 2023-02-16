package main

import (
	chat_service "github.com/gregb122/grpc-chat-protobufs/gen/go/chat_service/v1"
)

type InMemoryAuth struct {
	users []chat_service.UserInfo
}

type UserAuth interface {
	RegisterUser(req chat_service.RegisterUserRequest)
	LoginUser(req chat_service.LoginUserRequest)
	ListRegistredUsers() []chat_service.UserInfo
}
