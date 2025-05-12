package server

import "time"

type Transaction struct {
	ID           int       `gorm:"primaryKey;type:integer;" json:"id"`
	CurrencyCode uint16    `gorm:"type:integer;" json:"currency_code"`
	Amount       float64   `json:"amount"`
	BankName     string    `json:"bank_name"`
	Status       uint8     `gorm:"type:smallint;" json:"status"`
	CreatedAt    time.Time `gorm:"type:timestamp;" json:"created_at"`
}

type UpdateRequest struct {
	Status uint8 `json:"status"`
}
