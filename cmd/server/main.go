package main

import (
	"log"
	"net/http"

	"github.com/dgre1999/GolangCalculator/internal/api"
	"github.com/dgre1999/GolangCalculator/internal/calculator"
)

func main() {
	BasicCalc := calculator.NewBasic()
	RPNCalc := calculator.NewRPN()
	calcs := []calculator.Calculator{BasicCalc, RPNCalc}
	handler := api.NewHandler(calcs)

	http.HandleFunc("/api/v1/calc", handler.ComputeHandler)
	http.HandleFunc("/api/v1/history", handler.HistoryHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
