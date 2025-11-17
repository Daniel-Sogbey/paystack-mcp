package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Daniel-Sogbey/paystack-mcp-server/internals/paystack"
)

type application struct {
	Client *paystack.Client
}

type config struct {
	port    string
	baseUrl string
	apiKey  string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.port, "port", os.Getenv("PORT"), "mcp-server-port")
	flag.StringVar(&cfg.baseUrl, "baseUrl", os.Getenv("BASE_URL"), "paystack api base url")
	flag.StringVar(&cfg.apiKey, "apiKey", os.Getenv("PAYSTACK_API_KEY"), "paystack api key")

	flag.Parse()

	paystackClient := paystack.New(cfg.apiKey, cfg.baseUrl)

	app := &application{
		Client: paystackClient,
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.port),
		Handler: app.routes(),
	}

	log.Printf("paysack mcp server listening on port %s\n", cfg.port)
	log.Fatal(srv.ListenAndServe())
}
