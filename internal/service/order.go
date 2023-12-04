package service

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/sonochiwa/wb-level-0/internal/models"
	"github.com/sonochiwa/wb-level-0/internal/repository"
	"strconv"
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

	order := &models.Order{
		TrackNumber: gofakeit.Word(),
		Entry:       gofakeit.BuzzWord(),
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
			Amount:       gofakeit.Number(0, 1000),
			PaymentDt:    gofakeit.Number(0, 100000000),
			Bank:         gofakeit.Name(),
			DeliveryCost: gofakeit.Number(0, 3000),
			GoodsTotal:   gofakeit.Number(0, 400),
			CustomFee:    gofakeit.Number(0, 1000),
		},
		Locale:            gofakeit.Language(),
		InternalSignature: gofakeit.Word(),
		CustomerID:        gofakeit.UUID(),
		DeliveryService:   gofakeit.AppName(),
		ShardKey:          strconv.Itoa(gofakeit.Number(0, 1000)),
		SmID:              gofakeit.Number(0, 1000),
		DateCreated:       gofakeit.Date(),
		OofShard:          strconv.Itoa(gofakeit.Number(0, 1000)),
	}

	return s.repo.CreateOrder(*order)
}

func (s *OrderService) DeleteAllOrders() {
	s.repo.DeleteAllOrders()
}
