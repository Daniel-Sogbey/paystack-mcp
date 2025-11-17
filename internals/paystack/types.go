package paystack

import (
	"encoding/json"
	"time"
)

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
	Reference string
}

type VerifyPaymentResponse struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		ID              int64     `json:"id,omitempty"`
		Domain          string    `json:"domain,omitempty"`
		Status          string    `json:"status,omitempty"`
		Reference       string    `json:"reference,omitempty"`
		ReceiptNumber   any       `json:"receipt_number,omitempty"`
		Amount          int       `json:"amount,omitempty"`
		Message         any       `json:"message,omitempty"`
		GatewayResponse string    `json:"gateway_response,omitempty"`
		PaidAt          time.Time `json:"paid_at,omitempty"`
		CreatedAt       time.Time `json:"created_at,omitempty"`
		Channel         string    `json:"channel,omitempty"`
		Currency        string    `json:"currency,omitempty"`
		IPAddress       string    `json:"ip_address,omitempty"`
		Metadata        string    `json:"metadata,omitempty"`
		Log             struct {
			StartTime int   `json:"start_time,omitempty"`
			TimeSpent int   `json:"time_spent,omitempty"`
			Attempts  int   `json:"attempts,omitempty"`
			Errors    int   `json:"errors,omitempty"`
			Success   bool  `json:"success,omitempty"`
			Mobile    bool  `json:"mobile,omitempty"`
			Input     []any `json:"input,omitempty"`
			History   []struct {
				Type    string `json:"type,omitempty"`
				Message string `json:"message,omitempty"`
				Time    int    `json:"time,omitempty"`
			} `json:"history,omitempty"`
		} `json:"log,omitempty"`
		Fees          int `json:"fees,omitempty"`
		FeesSplit     any `json:"fees_split,omitempty"`
		Authorization struct {
			AuthorizationCode string `json:"authorization_code,omitempty"`
			Bin               string `json:"bin,omitempty"`
			Last4             string `json:"last4,omitempty"`
			ExpMonth          string `json:"exp_month,omitempty"`
			ExpYear           string `json:"exp_year,omitempty"`
			Channel           string `json:"channel,omitempty"`
			CardType          string `json:"card_type,omitempty"`
			Bank              string `json:"bank,omitempty"`
			CountryCode       string `json:"country_code,omitempty"`
			Brand             string `json:"brand,omitempty"`
			Reusable          bool   `json:"reusable,omitempty"`
			Signature         string `json:"signature,omitempty"`
			AccountName       any    `json:"account_name,omitempty"`
		} `json:"authorization,omitempty"`
		Customer struct {
			ID                       int    `json:"id,omitempty"`
			FirstName                any    `json:"first_name,omitempty"`
			LastName                 any    `json:"last_name,omitempty"`
			Email                    string `json:"email,omitempty"`
			CustomerCode             string `json:"customer_code,omitempty"`
			Phone                    any    `json:"phone,omitempty"`
			Metadata                 any    `json:"metadata,omitempty"`
			RiskAction               string `json:"risk_action,omitempty"`
			InternationalFormatPhone any    `json:"international_format_phone,omitempty"`
		} `json:"customer,omitempty"`
		Plan  any `json:"plan,omitempty"`
		Split struct {
		} `json:"split,omitempty"`
		OrderID            any       `json:"order_id,omitempty"`
		PaidAt0            time.Time `json:"paidAt,omitempty"`
		CreatedAt0         time.Time `json:"createdAt,omitempty"`
		RequestedAmount    int       `json:"requested_amount,omitempty"`
		PosTransactionData any       `json:"pos_transaction_data,omitempty"`
		Source             any       `json:"source,omitempty"`
		FeesBreakdown      any       `json:"fees_breakdown,omitempty"`
		Connect            any       `json:"connect,omitempty"`
		TransactionDate    time.Time `json:"transaction_date,omitempty"`
		PlanObject         struct {
		} `json:"plan_object,omitempty"`
		Subaccount struct {
		} `json:"subaccount,omitempty"`
	} `json:"data,omitempty"`
}
