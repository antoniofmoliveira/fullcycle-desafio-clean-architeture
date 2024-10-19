package web

import (
	"encoding/json"
	"net/http"

	"github.com/antoniofmoliveira/cleanarch/internal/usecase"
)

type WebOrderHandler struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewWebOrderHandler(
	CreateOrderUseCase usecase.CreateOrderUseCase,
) *WebOrderHandler {
	return &WebOrderHandler{
		CreateOrderUseCase: CreateOrderUseCase,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.CreateOrderUseCase.Execute(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
