package products

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {

	return &handler{service: s}

}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.GetProductsList(r.Context())
	if err != nil {
		slog.Error("Error marshalling data", "error", err)
		os.Exit(1)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(list); err != nil {
		slog.Error("Error marshalling data", "error", err)
		os.Exit(1)

	}
}
