package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhf0439/im-server/internal/service"
)

type UserHandler struct {
	users *service.UserService
}

func NewUserHandler(users *service.UserService) *UserHandler {
	return &UserHandler{users: users}
}

func (h *UserHandler) Me(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "user handler placeholder"})
}
