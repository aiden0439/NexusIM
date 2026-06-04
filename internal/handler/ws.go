package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhf0439/im-server/internal/ws"
)

type WSHandler struct {
	manager *ws.Manager
}

func NewWSHandler(manager *ws.Manager) *WSHandler {
	return &WSHandler{manager: manager}
}

func (h *WSHandler) Connect(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "websocket handler placeholder"})
}
