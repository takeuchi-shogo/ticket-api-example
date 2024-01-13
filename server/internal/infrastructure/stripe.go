package infrastructure

import (
	"log"
	"strconv"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/paymentsource"
	"github.com/stripe/stripe-go/v76/token"
	"github.com/takeuchi-shogo/ticket-api/config"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/gateways"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
)

type Stripe struct {
	PublicKey string
	SecretKey string
}

func NewStripe(c config.ServerConfig) gateways.Stripe {
	return &Stripe{
		PublicKey: c.StripePublicKey,
		SecretKey: c.StripeSecretKey,
	}
}

// 顧客の作成
func (s *Stripe) CreateCustomer(c *models.StripeCustomer) error {
	stripe.Key = s.SecretKey

	params := &stripe.CustomerParams{}
	result, err := customer.New(params)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

// 顧客の取得
func (s *Stripe) GetCustomer(customerID string) (*models.StripeCustomer, error) {
	stripe.Key = s.SecretKey

	params := &stripe.CustomerParams{}
	result, err := customer.Get(customerID, params)
	if err != nil {
		log.Println(err)
	}

	cus := &models.StripeCustomer{
		DefaultSource: result.DefaultSource,
	}

	return cus, nil
}

// クレカと顧客を紐付けた状態で作成
func (s *Stripe) CreateCreditCardAndCustomer(c *models.StripeCustomer, tk string) error {
	stripe.Key = s.SecretKey

	params := &stripe.CustomerParams{
		Name:  stripe.String(c.Name),
		Email: stripe.String(c.Email),
		Address: &stripe.AddressParams{
			City:       stripe.String(c.Address.City),
			Country:    stripe.String(c.Address.Country),
			Line1:      stripe.String(c.Address.Line1),
			Line2:      stripe.String(c.Address.Line2),
			PostalCode: stripe.String(c.Address.PostalCode),
			State:      stripe.String(c.Address.State),
		},
		Phone:  stripe.String(c.Phone),
		Source: stripe.String(c.Source),
	}
	result, err := customer.New(params)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

// クレジットカードを作成
func (s *Stripe) CreateCreditCard(customerID string) error {
	stripe.Key = s.SecretKey

	params := &stripe.PaymentSourceParams{
		Source: &stripe.PaymentSourceSourceParams{
			Token: stripe.String("tok_amex"),
		},
		Customer: stripe.String(customerID),
	}
	result, err := paymentsource.New(params)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

func (s *Stripe) GetCreditCardByCustomerID(customerID, defaultSource string) (*models.StripeCreditCards, error) {
	stripe.Key = s.SecretKey

	params := &stripe.PaymentSourceParams{
		Customer: &customerID,
	}
	result, err := paymentsource.Get(defaultSource, params)
	if err != nil {
		return &models.StripeCreditCards{}, err
	}
	return &models.StripeCreditCards{
		Number:   result.Card.Last4,
		Brand:    string(result.Card.Brand),
		ExpMonth: strconv.Itoa(int(result.Card.ExpMonth)),
		ExpYear:  strconv.Itoa(int(result.Card.ExpYear)),
	}, nil
}

// Customer にアタッチするカードのトークンを作成する（現在フロントで実装）
func (s *Stripe) CreateCreditCardToken(card *models.StripeCreditCards) (string, error) {
	stripe.Key = s.SecretKey

	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Name:     stripe.String(card.Name),
			Number:   stripe.String(card.Number),
			ExpMonth: stripe.String(card.ExpMonth),
			ExpYear:  stripe.String(card.ExpYear),
			CVC:      stripe.String(card.Cvc),
		},
	}
	result, err := token.New(params)
	if err != nil {
		return "", err
	}

	return result.ID, nil
}

// Stripeより発行されたトークンの取得
func (s *Stripe) GetToken(t string) error {
	stripe.Key = s.SecretKey

	params := &stripe.TokenParams{}
	_, err := token.Get(t, params)
	if err != nil {
		return err
	}
	return err
}
