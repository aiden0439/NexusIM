package repo

import (
	"github.com/aiden0439/NexusIM/internal/model"
	"gorm.io/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) Create(message *model.Message) error {
	return r.db.Create(message).Error
}
