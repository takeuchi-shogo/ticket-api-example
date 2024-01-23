package interactors

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type creditCardInteractor struct {
	db         usecase.DBUsecase
	creditCard usecase.CreditCardUsecase
	stripe     usecase.StripeUsecase
	user       usecase.UserUsecase
}

func NewCreditCardInteractor(
	db usecase.DBUsecase,
	creditCard usecase.CreditCardUsecase,
	stripe usecase.StripeUsecase,
	user usecase.UserUsecase,
) services.CreditCardService {
	return &creditCardInteractor{
		db:         db,
		creditCard: creditCard,
		stripe:     stripe,
		user:       user,
	}
}

func (c *creditCardInteractor) Get(userID int) (*models.CreditCardsReponse, *usecase.ResultStatus) {

	db, _ := c.db.Connect()

	creditCard, err := c.creditCard.FindByUserID(db, userID)
	if err != nil {
		return &models.CreditCardsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	customer, err := c.stripe.GetCustomer(creditCard.CustomerID)
	if err != nil {
		return &models.CreditCardsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	card, err := c.stripe.GetCreditCardByCustomerID(creditCard.CustomerID, customer.DefaultSource.ID)
	if err != nil {
		return &models.CreditCardsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return &models.CreditCardsReponse{
		Brand:    card.Brand,
		ExpMonth: card.ExpMonth,
		ExpYear:  card.ExpYear,
		Last4:    card.Number,
	}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (c creditCardInteractor) Create(userID int, token string) (*models.CreditCardsReponse, *usecase.ResultStatus) {

	db, _ := c.db.Connect()

	user, err := c.user.FindByID(db, userID)
	if err != nil {
		return &models.CreditCardsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	customer := &models.StripeCustomer{
		Name:   *user.DisplayName,
		Email:  user.Email,
		Source: token,
	}

	err = c.stripe.CreateCreditCardAndCustomer(customer, "")
	if err != nil {
		return &models.CreditCardsReponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return &models.CreditCardsReponse{}, usecase.NewResultStatus(http.StatusOK, nil)
}
