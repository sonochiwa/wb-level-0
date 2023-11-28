package order

type Repository interface {
	CreateOrder()
	GetByID()
	GetAllByNewsID()
}
