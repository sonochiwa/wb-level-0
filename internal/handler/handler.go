package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	mw "github.com/sonochiwa/wb-level-0/internal/middleware"
	"github.com/sonochiwa/wb-level-0/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.New(mw.GetCors()).Handler)
	r.Get("/", mainPage)
	r.Mount("/", h.orderRoutes())

	return r
}

func (h *Handler) orderRoutes() http.Handler {
	rg := chi.NewRouter()
	rg.Get("/orders", h.getAllOrders)
	rg.Get("/orders/{orderID}", h.getOrderByID)
	rg.Post("/orders", h.createOrder)
	rg.Delete("/orders", h.deleteAllOrders)

	return rg
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/index.html")
}
