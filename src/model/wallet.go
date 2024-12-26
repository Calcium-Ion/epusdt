package model

import (
	"time"
)

type WalletAddress struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Token     string    `gorm:"column:token;type:varchar(50);not null;index" json:"token"`
	Status    int       `gorm:"column:status;default:1;not null" json:"status"` // 1:启用 2:禁用
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (WalletAddress) TableName() string {
	return "wallet_address"
}
