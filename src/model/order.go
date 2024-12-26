package model

import (
	"time"
)

type Order struct {
	ID                 int       `gorm:"column:id;primaryKey;autoIncrement"`
	TradeID            string    `gorm:"column:trade_id;unique;type:varchar(32);not null" json:"trade_id"`
	OrderID            string    `gorm:"column:order_id;unique;type:varchar(32);not null" json:"order_id"`
	BlockTransactionID string    `gorm:"column:block_transaction_id;index;type:varchar(128)" json:"block_transaction_id"`
	ActualAmount       float64   `gorm:"column:actual_amount;type:decimal(19,4);not null" json:"actual_amount"`
	Amount             float64   `gorm:"column:amount;type:decimal(19,4);not null" json:"amount"`
	Token              string    `gorm:"column:token;type:varchar(50);not null" json:"token"`
	Status             int       `gorm:"column:status;default:1;not null" json:"status"` // 1：等待支付，2：支付成功，3：已过期
	NotifyURL          string    `gorm:"column:notify_url;type:varchar(128);not null" json:"notify_url"`
	RedirectURL        string    `gorm:"column:redirect_url;type:varchar(128)" json:"redirect_url"`
	CallbackNum        int       `gorm:"column:callback_num;default:0" json:"callback_num"`
	CallbackConfirm    int       `gorm:"column:callback_confirm;default:2" json:"callback_confirm"` // 1是 2否
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt          time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Order) TableName() string {
	return "orders"
}
