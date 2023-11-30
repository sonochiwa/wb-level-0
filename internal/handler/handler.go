package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	mw "github.com/sonochiwa/wb-level-0/internal/middleware"
	"github.com/sonochiwa/wb-level-0/internal/service"
	"net/http"
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
	r.Mount("/orders", orderRoutes())

	return r
}

func orderRoutes() http.Handler {
	rg := chi.NewRouter()
	rg.Get("/", getAllOrders)
	rg.Get("/{orderID}", getOrderByID)
	rg.Post("/", createOrder)

	return rg
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/index.html")
}
