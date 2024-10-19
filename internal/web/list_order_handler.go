package web

import (
	"encoding/json"
	"net/http"

	"github.com/antoniofmoliveira/cleanarch/internal/usecase"
)

type WebListOrderHandler struct {
	ListOrderUseCase usecase.ListOrderUseCase
}

func NewWebListOrderHandler(
	ListOrderUseCase usecase.ListOrderUseCase,
) *WebListOrderHandler {
	return &WebListOrderHandler{
		ListOrderUseCase: ListOrderUseCase,
	}
}

func (h *WebListOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	output, err := h.ListOrderUseCase.Execute()
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
