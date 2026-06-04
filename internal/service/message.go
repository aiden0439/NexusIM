package service

import "github.com/aiden0439/NexusIM/internal/repo"

type MessageService struct {
	messages *repo.MessageRepo
}

func NewMessageService(messages *repo.MessageRepo) *MessageService {
	return &MessageService{messages: messages}
}
