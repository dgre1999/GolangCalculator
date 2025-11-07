package main

import (
	"log"
	"net/http"

	"github.com/dgre1999/GolangCalculator/internal/api"
	"github.com/dgre1999/GolangCalculator/internal/calculator"
)

func main() {
	calc := calculator.New()
	handler := api.NewHandler(calc)

	http.HandleFunc("/api/v1/calc", handler.ComputeHandler)
	http.HandleFunc("/api/v1/history", handler.HistoryHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
