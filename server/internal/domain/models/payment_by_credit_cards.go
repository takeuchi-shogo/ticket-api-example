package models

import "errors"

type PaymentByCreditCards struct {
	ID               int `bun:",pk,autoincrement"`
	UserBookTicketID int
	PaymentID        string
	IsValid          bool
	ExpireAt         int64

	CreatedAt int64
	UpdatedAt int64
}

func (p *PaymentByCreditCards) Validate() error {
	if err := p.checkUserBookTicketID(); err != nil {
		return err
	}
	if err := p.checkPaymentID(); err != nil {
		return err
	}
	if err := p.checkExpireAt(); err != nil {
		return err
	}

	return nil
}

func (p *PaymentByCreditCards) checkUserBookTicketID() error {
	if p.UserBookTicketID == 0 {
		return errors.New("user_book_ticket_idを入力してください")
	}
	return nil
}

func (p *PaymentByCreditCards) checkPaymentID() error {
	if p.PaymentID == "" {
		return errors.New("payment_idを入力してください")
	}
	return nil
}

func (p *PaymentByCreditCards) checkExpireAt() error {
	if p.ExpireAt <= 0 {
		return errors.New("expire_atを入力してください")
	}
	return nil
}
