package api

import (
	"encoding/json"
	"net/http"

	"github.com/dgre1999/GolangCalculator/internal/calculator"
)

type Handler struct {
	calc *calculator.Calculator
}

func NewHandler(calc *calculator.Calculator) *Handler {
	return &Handler{calc: calc}
}

func (h *Handler) ComputeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		X  float64 `json:"x"`
		Y  float64 `json:"y"`
		Op string  `json:"op"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.calc.Compute(req.X, req.Y, req.Op)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"result": res})
}

func (h *Handler) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(h.calc.History())
}
