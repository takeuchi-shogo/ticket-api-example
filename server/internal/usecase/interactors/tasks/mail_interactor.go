package tasks

import (
	"log"

	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services/tasks"
)

type mailInteractor struct {
	db            usecase.DBUsecase
	registerEmail usecase.RegisterEmailUsecase
}

func NewMailInteractor(
	db usecase.DBUsecase,
	registerEmail usecase.RegisterEmailUsecase,
) tasks.MailService {
	return &mailInteractor{
		db:            db,
		registerEmail: registerEmail,
	}
}

func (m *mailInteractor) Send() error {
	// 新規登録の未送信メールがあれば送信する
	err := m.sendRegisterEmail()
	return err
}

func (m *mailInteractor) sendRegisterEmail() error {
	db, _ := m.db.Connect()

	registerEmails, _ := m.registerEmail.FindByNotSend(db)

	for _, registerEmail := range registerEmails {
		// ここにメール送信処理を記述する
		// 送信サービスの使用APIを使用予定
		log.Println(registerEmail)
	}

	return nil
}
