package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"tests/internal/orders"
)

type ordersService interface {
	OrderGet(ctx context.Context, UUID uuid.UUID) (*orders.Order, error)
	OrderCreate(ctx context.Context, entity *orders.Order) error
}

type HttpHandler struct {
	service ordersService
}

func NewHttpHandler(service ordersService) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (h *HttpHandler) Register(router *httprouter.Router) {
	router.GET("/api/v1/orders/:uuid", h.Get)
	router.POST("/api/v1/orders", h.Create)
}

func (h *HttpHandler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	UUID, err := uuid.Parse(p.ByName("uuid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}
	data, err := h.service.OrderGet(r.Context(), UUID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	d, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(d)
}

func (h *HttpHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var entity orders.Order
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}
	err = json.Unmarshal(body, &entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	err = h.service.OrderCreate(r.Context(), &entity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	w.WriteHeader(http.StatusOK)
}
