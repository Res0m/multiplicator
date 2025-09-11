package multiplier

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

type MultiplierHandler struct {
	MultiplierService *MultiplierService
}

func NewMultiplierHandler(router *mux.Router, handler *MultiplierHandler) {
	router.HandleFunc("/get", handler.GetMultiplier).Methods("GET")
}

func roundToDecimal(value float64, decimalPlaces int) float64 {
	multiplier := math.Pow(10, float64(decimalPlaces))
	return math.Round(value*multiplier) / multiplier
}

func (handler *MultiplierHandler) GetMultiplier(w http.ResponseWriter, r *http.Request) {
	var response Result
	multipl := roundToDecimal(handler.MultiplierService.GenerateMultiplierPareto(1.0, 10000.0), 2)
	response = Result{
		Result: multipl,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
