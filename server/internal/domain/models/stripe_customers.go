package models

import stripe "github.com/stripe/stripe-go/v76"

type StripeCustomer struct {
	ID      string
	Address struct {
		City       string `json:"city"`
		Country    string `json:"country"`
		Line1      string `json:"line_1"`
		Line2      string `json:"line_2"`
		PostalCode string `json:"postal_code"`
		State      string `json:"state"`
	}
	Description   string `json:"description"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Source        string
	DefaultSource *stripe.PaymentSource `json:"default_source"`
}
