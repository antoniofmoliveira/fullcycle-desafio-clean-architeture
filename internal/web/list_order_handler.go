package web

import (
	"encoding/json"
	"net/http"

	"github.com/antoniofmoliveira/cleanarch/internal/entity"
	"github.com/antoniofmoliveira/cleanarch/internal/usecase"
	"github.com/antoniofmoliveira/cleanarch/pkg/events"
)

type WebListOrderHandler struct {
	EventDispatcher  events.EventDispatcherInterface
	OrderRepository  entity.OrderRepositoryInterface
	OrderListedEvent events.EventInterface
}

func NewWebListOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderListedEvent events.EventInterface,
) *WebListOrderHandler {
	return &WebListOrderHandler{
		EventDispatcher:  EventDispatcher,
		OrderRepository:  OrderRepository,
		OrderListedEvent: OrderListedEvent,
	}
}

func (h *WebListOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listOrder := usecase.NewListOrderUseCase(h.OrderRepository, h.OrderListedEvent, h.EventDispatcher)
	output, err := listOrder.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
