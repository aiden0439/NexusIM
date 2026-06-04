package repo

import (
	"github.com/zhf0439/im-server/internal/model"
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
