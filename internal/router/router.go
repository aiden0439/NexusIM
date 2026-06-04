package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/aiden0439/NexusIM/internal/handler"
	"github.com/aiden0439/NexusIM/internal/repo"
	"github.com/aiden0439/NexusIM/internal/service"
	"github.com/aiden0439/NexusIM/internal/ws"
	"github.com/aiden0439/NexusIM/pkg/config"
	"github.com/aiden0439/NexusIM/pkg/kafka"
	"gorm.io/gorm"
)

type Dependencies struct {
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Kafka  *kafka.Client
}

func New(deps Dependencies) *gin.Engine {
	gin.SetMode(deps.Config.Server.Mode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	userRepo := repo.NewUserRepo(deps.DB)
	messageRepo := repo.NewMessageRepo(deps.DB)

	userService := service.NewUserService(userRepo)
	messageService := service.NewMessageService(messageRepo)

	wsManager := ws.NewManager()

	pingHandler := handler.NewPingHandler()
	userHandler := handler.NewUserHandler(userService)
	messageHandler := handler.NewMessageHandler(messageService)
	wsHandler := handler.NewWSHandler(wsManager)

	r.GET("/ping", pingHandler.Ping)

	api := r.Group("/api/v1")
	{
		api.GET("/users/me", userHandler.Me)
		api.POST("/messages", messageHandler.Send)
		api.GET("/ws", wsHandler.Connect)
	}

	return r
}
