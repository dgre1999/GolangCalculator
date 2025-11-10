package api

import (
	"encoding/json"
	"net/http"

	"github.com/dgre1999/GolangCalculator/internal/calculator"
)

type Handler struct {
	calcs map[string]calculator.Calculator
}

func NewHandler(calcsSlice []calculator.Calculator) *Handler {
	calcs := make(map[string]calculator.Calculator)
	for _, calc := range calcsSlice {
		switch c := calc.(type) {
		case *calculator.BasicCalculator:
			calcs["basic"] = c
		case *calculator.RPNCalculator:
			calcs["rpn"] = c
		}
	}
	return &Handler{calcs: calcs}
}

func (h *Handler) ComputeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Type       string `json:"type"`
		Expression string `json:"expression"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Type != "basic" && req.Type != "rpn" {
		http.Error(w, "invalid calculator type", http.StatusBadRequest)
		return
	}

	calc := h.calcs[req.Type]
	res, err := calc.Compute(req.Expression)
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
	var combinedHistory []string
	for _, calc := range h.calcs {
		combinedHistory = append(combinedHistory, calc.History()...)
	}

	json.NewEncoder(w).Encode(combinedHistory)
}
