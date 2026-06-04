package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhf0439/im-server/internal/service"
)

type MessageHandler struct {
	messages *service.MessageService
}

func NewMessageHandler(messages *service.MessageService) *MessageHandler {
	return &MessageHandler{messages: messages}
}

func (h *MessageHandler) Send(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "message handler placeholder"})
}
