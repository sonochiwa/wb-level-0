package models

import (
	"encoding/json"
	"errors"
	"time"
)

func (m *Delivery) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}

func (m *Payment) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}

type Order struct {
	OrderUID          *string       `json:"order_uid" db:"order_uid"`
	TrackNumber       *string       `json:"track_number" db:"track_number"`
	Entry             *string       `json:"entry" db:"entry"`
	Delivery          *Delivery     `json:"delivery" db:"delivery"`
	Payment           *Payment      `json:"payment" db:"payment"`
	Items             *[]OrderItems `json:"items" db:"items"`
	Locale            *string       `json:"locale" db:"locale"`
	InternalSignature *string       `json:"internal_signature" db:"internal_signature"`
	CustomerID        *string       `json:"customer_id" db:"customer_id"`
	DeliveryService   *string       `json:"delivery_service" db:"delivery_service"`
	ShardKey          *string       `json:"shard_key" db:"shard_key"`
	SmID              *uint64       `json:"sm_id" db:"sm_id"`
	DateCreated       *time.Time    `json:"date_created" db:"date_created"`
	OofShard          *string       `json:"oof_shard" db:"oof_shard"`
}

type Delivery struct {
	Name    *string `json:"name" db:"name"`
	Phone   *string `json:"phone" db:"phone"`
	Zip     *string `json:"zip" db:"zip"`
	City    *string `json:"city" db:"city"`
	Address *string `json:"address" db:"address"`
	Region  *string `json:"region" db:"region"`
	Email   *string `json:"email" db:"email"`
}

type Payment struct {
	Transaction  *string `json:"transaction" db:"transaction"`
	RequestID    *string `json:"request_id" db:"request_id"`
	Currency     *string `json:"currency" db:"currency"`
	Provider     *string `json:"provider" db:"provider"`
	Amount       *uint64 `json:"amount" db:"amount"`
	PaymentDt    *uint64 `json:"payment_dt" db:"payment_dt"`
	Bank         *string `json:"bank" db:"bank"`
	DeliveryCost *uint64 `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   *uint64 `json:"goods_total" db:"goods_total"`
	CustomFee    *uint64 `json:"custom_fee" db:"custom_fee"`
}
