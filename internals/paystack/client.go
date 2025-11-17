package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Daniel-Sogbey/paystack-mcp-server/internals/requester"
)

type Client struct {
	APIKey  string
	BaseUrl string
}

func New(apiKey, baseUrl string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseUrl: baseUrl,
	}
}

func (c *Client) InitializePayment(req InitializePaymentRequest) (*InitializePaymentResponse, error) {
	url := fmt.Sprintf("%s/transaction/initialize", c.BaseUrl)

	header := http.Header{}
	header.Set("Content-Type", "application/json")
	header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := requester.Requester[InitializePaymentResponse](http.MethodPost, url, header, nil, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) VerifyPayment(req VerifyPaymentRequest) (*VerifyPaymentResponse, error) {
	return nil, nil
}
