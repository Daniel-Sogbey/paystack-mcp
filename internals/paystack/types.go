package paystack

import "encoding/json"

type InitializePaymentRequest struct {
	Email     string          `json:"email"`
	Amount    int             `json:"amount"`
	Reference string          `json:"reference"`
	Currency  string          `json:"currency"`
	Metadata  json.RawMessage `json:"metadata"`
}

type InitializePaymentResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationUrl string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
		Reference        string `json:"reference"`
	}
}

type VerifyPaymentRequest struct {
}

type VerifyPaymentResponse struct {
}
