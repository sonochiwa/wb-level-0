package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	gf "github.com/brianvoe/gofakeit/v6"
	"github.com/sonochiwa/wb-level-0/internal/clients/stan"
	mc "github.com/sonochiwa/wb-level-0/internal/memcache"
	"github.com/sonochiwa/wb-level-0/internal/models"
	"github.com/sonochiwa/wb-level-0/internal/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetAllOrders() ([]models.OrderID, error) {
	return s.repo.GetAllOrders()
}

func (s *OrderService) GetOrderById(orderID string) (models.Order, error) {
	var wanted string
	var jsonWander models.Order
	mc.MC.Get(context.TODO(), orderID, &wanted)

	err := json.Unmarshal([]byte(wanted), &jsonWander)
	if err != nil {
		fmt.Println(err)
	}

	return jsonWander, nil
}

func (s *OrderService) CreateOrder() error {
	trackNumber := gf.Word()
	items := make([]models.Item, gf.Number(1, 3))
	for item := range items {
		items[item] = models.Item{
			ChrtID:      gf.Number(0, 65535),
			TrackNumber: trackNumber,
			Price:       gf.Number(0, 65535),
			Rid:         gf.UUID(),
			Name:        gf.Name(),
			Sale:        gf.Number(0, 90),
			Size:        gf.Number(32, 62),
			TotalPrice:  gf.Number(0, 65535),
			NmID:        gf.UUID(),
			Brand:       gf.Name(),
			Status:      gf.Number(200, 202),
		}
	}
	order := &models.Order{
		OrderUID:    gf.UUID(),
		TrackNumber: trackNumber,
		Entry:       gf.Word(),
		Delivery: models.Delivery{
			Name:    gf.Name(),
			Phone:   gf.Phone(),
			Zip:     gf.Zip(),
			City:    gf.City(),
			Address: gf.Address().Address,
			Region:  gf.Country(),
			Email:   gf.Email(),
		},
		Payment: models.Payment{
			Transaction:  gf.UUID(),
			RequestID:    gf.UUID(),
			Currency:     gf.Currency().Short,
			Provider:     gf.Username(),
			Amount:       gf.Number(0, 65535),
			PaymentDt:    gf.Number(0, 65535),
			Bank:         gf.Name(),
			DeliveryCost: gf.Number(0, 65535),
			GoodsTotal:   gf.Number(0, 65535),
			CustomFee:    gf.Number(0, 65535),
		},
		Items:             items,
		Locale:            gf.Language(),
		InternalSignature: gf.Word(),
		CustomerID:        gf.UUID(),
		DeliveryService:   gf.AppName(),
		ShardKey:          strconv.Itoa(gf.Number(0, 65535)),
		SmID:              gf.Number(0, 65535),
		DateCreated:       gf.Date(),
		OofShard:          strconv.Itoa(gf.Number(0, 65535)),
	}

	message, err := json.Marshal(order)
	if err != nil {
		return err
	}

	// TODO: validate data before public to channel
	err = stan.PublishMessage(message)
	if err != nil {
		return err
	}

	//s.repo.CreateOrder(*order)

	return nil
}

func (s *OrderService) DeleteAllOrders() error {
	return s.repo.DeleteAllOrders()
}
