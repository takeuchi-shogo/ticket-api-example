package models

import (
	"errors"

	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

type Users struct {
	ID          uint64  `json:"id" bun:",pk,autoincrement"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	DisplayName *string `json:"display_name"`
	ScreenName  string  `json:"screen_name"`
	Email       string  `json:"email"`
	Tel         string  `json:"tel"`
	Password    string  `json:"password"`
	PostCode    string  `json:"post_code"`
	Prefecture  string  `json:"prefecture"`
	City        string  `json:"city"`

	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"deleted_at"`
}

type UsersResponse struct {
	ID          uint64  `json:"id"`
	DisplayName *string `json:"display_name"`
	ScreenName  string  `json:"screen_name"`
	Email       string  `json:"email"`
}

type MeInteractorResponse struct {
	User  *UsersResponse    `json:"user"`
	Token *token.TokenPairs `json:"token"`
}

func (u *Users) Validate() error {
	if err := u.validateFirstName(); err != nil {
		return err
	}
	if err := u.validateLastName(); err != nil {
		return err
	}
	if err := u.validateScreenName(); err != nil {
		return err
	}
	if err := u.validateEmail(); err != nil {
		return err
	}
	if err := u.validatePassword(); err != nil {
		return err
	}
	return nil
}

func (u *Users) validateFirstName() error {
	if u.FirstName == "" {
		return errors.New("姓を入力してください")
	}
	return nil
}

func (u *Users) validateLastName() error {
	if u.LastName == "" {
		return errors.New("名を入力してください")
	}
	return nil
}

func (u *Users) validateScreenName() error {
	if u.ScreenName == "" {
		return errors.New("登録エラー")
	}
	return nil
}

func (u *Users) validateEmail() error {
	if u.Email == "" {
		return errors.New("メールアドレスを入力してください")
	}
	return nil
}

func (u *Users) validatePassword() error {
	if u.Password == "" {
		return errors.New("パスワードを入力してください")
	}
	return nil
}

func (u *Users) BuildForGet() *UsersResponse {
	return &UsersResponse{
		ID:          u.ID,
		DisplayName: u.DisplayName,
		ScreenName:  u.ScreenName,
		Email:       u.Email,
	}
}
