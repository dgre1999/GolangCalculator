package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/dgre1999/GolangCalculator/internal/calculator"
)

var users = []string{
	"daniel",
	"redia",
}
var passmap = map[string]string{
	"daniel": "bc46a8a0c0be4569980c46022ed56596eca1b2c50b3b2fb57c88383519c92c7c",
	"redia":  "2bfc003061a6a8f15f0705fbd4151ac167d405595c83e0548c872fbee69234f0",
}

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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	(*w).Header().Set("Vary", "Origin")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func (h *Handler) ComputeHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Type       string `json:"type"`
		Expression string `json:"expression"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !validateUserAndPass(req.Username, req.Password) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
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
	enableCors(&w)
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

func getSHA256Hash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func validateUserAndPass(user, pass string) bool {
	if slices.Contains(users, user) && passmap[user] == getSHA256Hash(pass) {
		return true
	}
	return false
}
