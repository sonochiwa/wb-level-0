package service

import (
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
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
	return s.repo.GetOrderById(orderID)
}

func (s *OrderService) CreateOrder() (string, error) {
	trackNumber := gofakeit.Word()

	items := make([]models.Item, gofakeit.Number(1, 3))

	for k, _ := range items {
		items[k] = models.Item{
			ChrtID:      gofakeit.Number(0, 65535),
			TrackNumber: trackNumber,
			Price:       gofakeit.Number(0, 65535),
			Rid:         gofakeit.UUID(),
			Name:        gofakeit.Name(),
			Sale:        gofakeit.Number(0, 90),
			Size:        gofakeit.Number(32, 62),
			TotalPrice:  gofakeit.Number(0, 65535),
			NmID:        gofakeit.UUID(),
			Brand:       gofakeit.Name(),
			Status:      gofakeit.Number(200, 202),
		}
	}

	order := &models.Order{
		TrackNumber: trackNumber,
		Entry:       gofakeit.Word(),
		Delivery: models.Delivery{
			Name:    gofakeit.Name(),
			Phone:   gofakeit.Phone(),
			Zip:     gofakeit.Zip(),
			City:    gofakeit.City(),
			Address: gofakeit.Address().Address,
			Region:  gofakeit.Country(),
			Email:   gofakeit.Email(),
		},
		Payment: models.Payment{
			Transaction:  gofakeit.UUID(),
			RequestID:    gofakeit.UUID(),
			Currency:     gofakeit.Currency().Short,
			Provider:     gofakeit.Username(),
			Amount:       gofakeit.Number(0, 65535),
			PaymentDt:    gofakeit.Number(0, 65535),
			Bank:         gofakeit.Name(),
			DeliveryCost: gofakeit.Number(0, 65535),
			GoodsTotal:   gofakeit.Number(0, 65535),
			CustomFee:    gofakeit.Number(0, 65535),
		},
		Items:             items,
		Locale:            gofakeit.Language(),
		InternalSignature: gofakeit.Word(),
		CustomerID:        gofakeit.UUID(),
		DeliveryService:   gofakeit.AppName(),
		ShardKey:          strconv.Itoa(gofakeit.Number(0, 65535)),
		SmID:              gofakeit.Number(0, 65535),
		DateCreated:       gofakeit.Date(),
		OofShard:          strconv.Itoa(gofakeit.Number(0, 65535)),
	}

	return s.repo.CreateOrder(*order)
}

func (s *OrderService) DeleteAllOrders() {
	s.repo.DeleteAllOrders()
}
