package model

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"size:64;uniqueIndex;not null" json:"username"`
	Nickname  string    `gorm:"size:64" json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Conversation struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Type      string    `gorm:"size:32;not null" json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ConversationMember struct {
	ID             uint64    `gorm:"primaryKey" json:"id"`
	ConversationID uint64    `gorm:"not null;uniqueIndex:uk_conversation_members_conversation_user;index:idx_conversation_members_user_conversation,priority:2" json:"conversation_id"`
	UserID         uint64    `gorm:"not null;uniqueIndex:uk_conversation_members_conversation_user;index:idx_conversation_members_user_conversation,priority:1" json:"user_id"`
	JoinedAt       time.Time `json:"joined_at"`
}

type Message struct {
	ID              uint64    `gorm:"primaryKey" json:"id"`
	ConversationID  uint64    `gorm:"not null;uniqueIndex:uk_messages_conversation_seq;index:idx_messages_conversation_created,priority:1" json:"conversation_id"`
	SenderID        uint64    `gorm:"not null;uniqueIndex:uk_messages_sender_client_message" json:"sender_id"`
	ClientMessageID string    `gorm:"size:64;not null;uniqueIndex:uk_messages_sender_client_message" json:"client_message_id"`
	Seq             uint64    `gorm:"not null;uniqueIndex:uk_messages_conversation_seq" json:"seq"`
	Type            string    `gorm:"size:32;not null;default:text" json:"type"`
	Content         string    `gorm:"type:text;not null" json:"content"`
	Status          string    `gorm:"size:32;not null;default:sent" json:"status"`
	CreatedAt       time.Time `gorm:"index:idx_messages_conversation_created,priority:2" json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
