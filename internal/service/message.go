package service

import "github.com/zhf0439/im-server/internal/repo"

type MessageService struct {
	messages *repo.MessageRepo
}

func NewMessageService(messages *repo.MessageRepo) *MessageService {
	return &MessageService{messages: messages}
}
