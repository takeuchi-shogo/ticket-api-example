package models

type StripeCreditCards struct {
	Name     string `json:"name"`
	Number   string `json:"number"`
	Brand    string `json:"brand"`
	ExpMonth string `json:"exp_month"`
	ExpYear  string `json:"exp_year"`
	Cvc      string `json:"cvc"`
}
