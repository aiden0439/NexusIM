package repo

import (
	"github.com/aiden0439/NexusIM/internal/model"
	"gorm.io/gorm"
)

type ConversationRepo struct {
	db *gorm.DB
}

func NewConversationRepo(db *gorm.DB) *ConversationRepo {
	return &ConversationRepo{db: db}
}

func (r *ConversationRepo) FindByID(id uint64) (*model.Conversation, error) {
	var conversation model.Conversation
	if err := r.db.First(&conversation, id).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}
