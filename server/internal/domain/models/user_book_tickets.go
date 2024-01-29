package models

import "errors"

type UserBookTickets struct {
	ID              int    `json:"id" bun:",pk,autoincrement"`
	BookID          string `json:"book_id"`
	UserID          int    `json:"user_id"`
	EventID         int    `json:"event_id"`
	TicketID        int    `json:"ticket_id"`
	TicketItemID    int    `json:"ticket_item_id"`
	PaymentMethod   string `json:"payment_method"`
	NumberOfTickets int    `json:"number_of_tickets"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type UserBookTicketsReponse struct {
	ID              int    `json:"id"`
	BookID          string `json:"book_id"`
	UserID          int    `json:"user_id"`
	EventID         int    `json:"event_id"`
	TicketID        int    `json:"ticket_id"`
	TicketItemID    int    `json:"ticket_item_id"`
	PaymentMethod   string `json:"payment_method"`
	NumberOfTickets int    `json:"number_of_tickets"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	Event          *EventsReponse            `json:"event"`
	Ticket         *TicketsResponse          `json:"ticket"`
	UserHasTickets []*UserHasTicketsResponse `json:"user_has_tickets"`
}

func (u *UserBookTickets) Validate() error {

	if err := u.checkBookID(); err != nil {
		return err
	}
	if err := u.checkUserID(); err != nil {
		return err
	}
	if err := u.checkEventID(); err != nil {
		return err
	}
	if err := u.checkTicketID(); err != nil {
		return err
	}
	if err := u.checkTicketItemID(); err != nil {
		return err
	}
	if err := u.checkPaymentMethod(); err != nil {
		return err
	}
	if err := u.checkNumberOfTickets(); err != nil {
		return err
	}
	return nil
}

func (u *UserBookTickets) checkBookID() error {
	if u.BookID == "" {
		return errors.New("書籍IDを入力してください")
	}
	return nil
}

func (u *UserBookTickets) checkUserID() error {
	if u.UserID == 0 {
		return errors.New("ユーザIDを入力してください")
	}
	return nil
}

func (u *UserBookTickets) checkEventID() error {
	if u.EventID == 0 {
		return errors.New("イベントIDを入力してください")
	}
	return nil
}

func (u *UserBookTickets) checkTicketID() error {
	if u.TicketID == 0 {
		return errors.New("チケットIDを入力してください")
	}
	return nil
}

func (u *UserBookTickets) checkTicketItemID() error {
	if u.TicketItemID == 0 {
		return errors.New("チケットアイテムIDを入力してください")
	}
	return nil
}

func (u *UserBookTickets) checkPaymentMethod() error {
	if u.PaymentMethod == "" {
		return errors.New("支払い方法を入力してください")
	}
	return nil
}

func (u *UserBookTickets) checkNumberOfTickets() error {
	if u.NumberOfTickets == 0 {
		return errors.New("チケット枚数を入力してください")
	}
	return nil
}

func (u *UserBookTickets) BuildForGet() *UserBookTicketsReponse {
	return &UserBookTicketsReponse{
		ID:              u.ID,
		BookID:          u.BookID,
		UserID:          u.UserID,
		EventID:         u.EventID,
		TicketID:        u.TicketID,
		TicketItemID:    u.TicketItemID,
		PaymentMethod:   u.PaymentMethod,
		NumberOfTickets: u.NumberOfTickets,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
	}
}
