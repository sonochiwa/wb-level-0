package models

import (
	"encoding/json"
	"errors"
	"time"
)

type ItemsScanner []Item

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

func (m *Item) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}

func (items *ItemsScanner) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, items)
}

type OrderID struct {
	OrderUID string `json:"order_uid" db:"order_uid"`
}

type Order struct {
	OrderUID          string       `json:"order_uid" db:"order_uid"`
	TrackNumber       string       `json:"track_number" db:"track_number"`
	Entry             string       `json:"entry" db:"entry"`
	Delivery          Delivery     `json:"delivery" db:"delivery"`
	Payment           Payment      `json:"payment" db:"payment"`
	Items             ItemsScanner `json:"items" db:"items"`
	Locale            string       `json:"locale" db:"locale"`
	InternalSignature string       `json:"internal_signature" db:"internal_signature"`
	CustomerID        string       `json:"customer_id" db:"customer_id"`
	DeliveryService   string       `json:"delivery_service" db:"delivery_service"`
	ShardKey          string       `json:"shard_key" db:"shard_key"`
	SmID              int          `json:"sm_id" db:"sm_id"`
	DateCreated       time.Time    `json:"date_created" db:"date_created"`
	OofShard          string       `json:"oof_shard" db:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Zip     string `json:"zip" db:"zip"`
	City    string `json:"city" db:"city"`
	Address string `json:"address" db:"address"`
	Region  string `json:"region" db:"region"`
	Email   string `json:"email" db:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction" db:"transaction"`
	RequestID    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int    `json:"amount" db:"amount"`
	PaymentDt    int    `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        int    `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        string `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
