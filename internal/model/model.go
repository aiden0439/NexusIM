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

type Message struct {
	ID             uint64    `gorm:"primaryKey" json:"id"`
	ConversationID uint64    `gorm:"index;not null" json:"conversation_id"`
	SenderID       uint64    `gorm:"index;not null" json:"sender_id"`
	Seq            uint64    `gorm:"index;not null" json:"seq"`
	Content        string    `gorm:"type:text;not null" json:"content"`
	Status         string    `gorm:"size:32;not null" json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
