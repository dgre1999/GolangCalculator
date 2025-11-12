package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/dgre1999/GolangCalculator/internal/calculator"
)

// Valid usernames
var users = []string{
	"daniel",
	"redia",
}

// Map of valid usernames and the hashed passwords
var passmap = map[string]string{
	"daniel": "bc46a8a0c0be4569980c46022ed56596eca1b2c50b3b2fb57c88383519c92c7c",
	"redia":  "2bfc003061a6a8f15f0705fbd4151ac167d405595c83e0548c872fbee69234f0",
}

// Contains a map of calculator types to their instances, since we want to be able to have a calculator of each type
type Handler struct {
	calcs map[string]calculator.Calculator
}

// Creates a handler with the corresponding calculators in the map
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

// Enables CORS for the specified origin. If not enabled, gives CORS errors in browser
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "https://fir-testing-323bd.web.app/")
	(*w).Header().Set("Vary", "Origin")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// Handles the HTTP requests to the compute function
func (h *Handler) ComputeHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// Handle preflight OPTIONS request, if not handled, gives CORS errors in browser
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	// If not a POST request, stop
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// The JSON format the API wants, with authentication details first, followed by the type of calculator to use, and lastly the expression to compute
	var req struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Type       string `json:"type"`
		Expression string `json:"expression"`
	}
	// Decode the JSON request body into the req struct while checking for errors
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Authenticate
	if !validateUserAndPass(req.Username, req.Password) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// Check if the requested calculator type is valid (maybe more elegant solution than if-else?)
	if req.Type != "basic" && req.Type != "rpn" {
		http.Error(w, "invalid calculator type", http.StatusBadRequest)
		return
	}
	// Compute the expression using the requested calculator
	calc := h.calcs[req.Type]
	res, err := calc.Compute(req.Expression)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Return the result as JSON
	json.NewEncoder(w).Encode(map[string]any{"result": res})
}

// Handles the HTTP requests for the History function
func (h *Handler) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// Checks if GET, else stop
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Combine history from all calculators (need to find a way to sync, but not sure how to implement that yet)
	var combinedHistory []string
	for _, calc := range h.calcs {
		combinedHistory = append(combinedHistory, calc.History()...)
	}
	// Return the combined history as JSON
	json.NewEncoder(w).Encode(combinedHistory)
}

// Gets the sha256 hash of the input string, used to validate username and password
func getSHA256Hash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// Checks if the username exists, and if the given password matches the given user
func validateUserAndPass(user, pass string) bool {
	if slices.Contains(users, user) && passmap[user] == getSHA256Hash(pass) {
		return true
	}
	return false
}
