package transactions

import (
	"errors"
	"mini-project/modules/products"
	"time"
)

type Transaction struct {
	ID        int               `gorm:"primaryKey" json:"id"`
	Timestamp time.Time         `json:"timestamp"`
	Total     int               `gorm:"not null" json:"total"`
	Items     []TransactionItem `json:"items"`
}

type TransactionItem struct {
	ID            int               `gorm:"primaryKey" json:"id"`
	TransactionID uint              `gorm:"not null;foreignKey:TransactionID"`
	ProductID     uint              `gorm:"not null;foreignKey:ProductID"`
	Quantity      int               `gorm:"not null" json:"quantity"`
	Price         int               `gorm:"not null" json:"price"`
	Product       *products.Product `json:"product"`
}

type CreateItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CreateTransactionRequest struct {
	Items []CreateItemRequest
}

var (
	ErrProductIdNotFound = errors.New("Product id not found")
	ErrStockNotEnough = errors.New("Stock not enough")
	ErrPoductHasBeenRemoved = errors.New("Product has been removed")
)
