package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	mux := httprouter.New()

	mux.HandlerFunc(http.MethodGet, "/.well-known/mcp-manifest", app.manifestHandler)
	mux.HandlerFunc(http.MethodPost, "/mcp/paystack/initialize", app.initializePaymentHandler)
	mux.HandlerFunc(http.MethodGet, "/mcp/paystack/verify", app.verifyPayment)

	return mux
}
