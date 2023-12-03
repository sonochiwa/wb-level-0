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
	trackNumber := gofakeit.Word()
	entry := gofakeit.BuzzWord()

	name := gofakeit.Name()
	phone := gofakeit.Phone()
	zip := gofakeit.Zip()
	city := gofakeit.City()
	address := gofakeit.Address().Address
	region := gofakeit.Country()
	email := gofakeit.Email()

	delivery := models.Delivery{
		Name:    &name,
		Phone:   &phone,
		Zip:     &zip,
		City:    &city,
		Address: &address,
		Region:  &region,
		Email:   &email,
	}

	requestId := gofakeit.UUID()
	currency := gofakeit.Currency().Short
	provider := gofakeit.Username()
	amount := gofakeit.Number(0, 1000)
	paymentDt := gofakeit.Number(0, 100000000)
	bank := gofakeit.Name()
	deliveryCost := gofakeit.Number(0, 3000)
	goodsTotal := gofakeit.Number(0, 400)
	customFee := gofakeit.Number(0, 1000)

	payment := models.Payment{
		RequestID:    &requestId,
		Currency:     &currency,
		Provider:     &provider,
		Amount:       &amount,
		PaymentDt:    &paymentDt,
		Bank:         &bank,
		DeliveryCost: &deliveryCost,
		GoodsTotal:   &goodsTotal,
		CustomFee:    &customFee,
	}

	locale := gofakeit.Language()
	internalSignature := gofakeit.Word()
	customerId := gofakeit.UUID()
	deliveryService := gofakeit.AppName()
	shardKey := strconv.Itoa(gofakeit.Number(0, 1000))
	smID := gofakeit.Number(0, 1000)
	dateCreated := gofakeit.Date()
	oofShard := strconv.Itoa(gofakeit.Number(0, 1000))

	order := models.Order{
		TrackNumber:       &trackNumber,
		Entry:             &entry,
		Delivery:          &delivery,
		Payment:           &payment,
		Locale:            &locale,
		InternalSignature: &internalSignature,
		CustomerID:        &customerId,
		DeliveryService:   &deliveryService,
		ShardKey:          &shardKey,
		SmID:              &smID,
		DateCreated:       &dateCreated,
		OofShard:          &oofShard,
	}

	return s.repo.CreateOrder(order)
}

func (s *OrderService) DeleteAllOrders() {
	s.repo.DeleteAllOrders()
}
