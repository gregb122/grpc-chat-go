package main

import (
	"errors"
	"fmt"

	chat_service "github.com/gregb122/grpc-chat-protobufs/gen/go/chat_service/v1"
)

type Message struct {
	id  string
	msg chat_service.Message
}
type MessageHandler interface {
	AddMessageToQueue(toU, msg string) error
	GetElemsFromQueue(isToSend bool, getAll bool) ([]Message, error)
	StoreAndDeleteSentMessages(msgs []Message) error
}

type userQueues struct {
	to_user       string
	to_send_queue []chat_service.Message
	sent_sotarage []chat_service.Message
}

type inMemorToUsers struct {
	toUsers map[string]*userQueues
}

func (que *userQueues) init(to_user string) {
	fmt.Printf("Initialization of messages queues for %s\n", to_user)
	que.to_user = to_user
	que.to_send_queue = make([]chat_service.Message, 0)
	que.sent_sotarage = make([]chat_service.Message, 0)
}

func (u *InMemorToUsers) init() {
	fmt.Println("Initialization of memory of users")

	u.toUsers = make(map[string]*userQueues)
}

func (que *InMemorToUsers) AddMessageToQueue(
	to_user string,
	msg chat_service.Message,
	isToSend bool) error {

	_, ok := que.toUsers[to_user]
	if !ok {
		return errors.New(fmt.Sprintf("Destination user <%s? not found", to_user))
	}
	if isToSend {
		que.toUsers[to_user].to_send_queue =
			append(que.toUsers[to_user].to_send_queue, msg)
	} else {
		que.toUsers[to_user].sent_sotarage =
			append(que.toUsers[to_user].sent_sotarage, msg)
	}

	return nil
}

func (que *InMemorToUsers) GetElemsFromQueue(to_user string, isToSend, getAll bool) {
	_, ok := que.toUsers[to_user]
	if !ok {
		return errors.New(fmt.Sprintf("Destination user <%s? not found", to_user))
	}
}
