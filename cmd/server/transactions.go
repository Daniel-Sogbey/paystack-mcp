package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Daniel-Sogbey/paystack-mcp-server/internals/paystack"
	"github.com/Daniel-Sogbey/paystack-mcp-server/internals/types"
)

func (app *application) manifestHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("./.well-known/manifest/manifest.json")
	if err != nil {
		log.Println("Error Log1:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var manifest types.Manifest
	err = json.Unmarshal(b, &manifest)
	if err != nil {
		log.Println("Error Log2:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(manifest)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (app *application) initializePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var req paystack.InitializePaymentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	resp, err := app.Client.InitializePayment(req)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (app *application) verifyPayment(w http.ResponseWriter, r *http.Request) {

}
