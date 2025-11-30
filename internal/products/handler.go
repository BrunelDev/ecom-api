package products

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {

	return &handler{service: s}

}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"message": "Django is mid as fuck!!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(res)
	// data, err := json.Marshal(res)
	if err != nil {
		slog.Error("Error marshalling data", "error", err)
	}

	// w.Write(data)

}
