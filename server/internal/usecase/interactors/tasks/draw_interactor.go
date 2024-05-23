package tasks

import (
	"log"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/enums"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services/tasks"
	"github.com/takeuchi-shogo/ticket-api/pkg/random"
	"github.com/uptrace/bun"
)

type drawInteractor struct {
	db             usecase.DBUsecase
	ticket         usecase.TicketUsecase
	userBookTicket usecase.UserBookTicketUsecase
	userHasTicket  usecase.UserHasTicketUsecase
	ticketStatus   *enums.UserHasTicketStatus
}

func NewDrawInteractor(
	db usecase.DBUsecase,
	ticket usecase.TicketUsecase,
	userBookTicket usecase.UserBookTicketUsecase,
	userHasTicket usecase.UserHasTicketUsecase,
) tasks.DrawService {
	return &drawInteractor{
		db:             db,
		ticket:         ticket,
		userBookTicket: userBookTicket,
		userHasTicket:  userHasTicket,
		ticketStatus:   enums.NewUserHasTicketStatus(),
	}
}

// 抽選を開始
func (d *drawInteractor) StartDrawing() error {

	db, _ := d.db.Connect()

	// 抽選を行う販売スケジュールの取得
	tickets, err := d.ticket.FindByDrawingAt(db, time.Now().Unix())
	if err != nil {
		return nil
	}

	// スケジュール毎に抽選を行う
	for _, ticket := range tickets {
		err = d.drawingByTicket(db, ticket)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (d *drawInteractor) drawingByTicket(db bun.IDB, ticket *models.Tickets) error {
	// 抽選中のチケットの取得
	userBookTickets, err := d.userBookTicket.FindByTicketID(db, ticket.ID)
	if err != nil {
		return err
	}

	// 当選チケット枚数を管理
	numberOfWinningTickets := 0
	// 申し込み枚数
	numberOfAppliedTicketCnt := 0
	// 設定されている当選予定のチケットカウント
	numberOfWinningTicketCnt := ticket.NumberOfWinningTickets
	// 落選する予定のチケットの算出
	numberOfNotWonTicketCnt := 0
	// 当選したチケットはtrueにしてそれ以外はfalseで管理するマップ
	winningUserBookTickets := map[int]bool{}

	userBookTicketsTemp := map[int][]*models.UserHasTickets{}
	for _, userBookTicket := range userBookTickets {
		userHasTickets, err := d.userHasTicket.FindByUserBookTicketID(db, userBookTicket.ID)
		if err != nil {
			continue
		}
		numberOfAppliedTicketCnt += userBookTicket.NumberOfTickets
		userBookTicketsTemp[userBookTicket.ID] = userHasTickets
		winningUserBookTickets[userBookTicket.ID] = false
	}

	// 当選処理
	for {
		// 当選枚数に達すると当選処理を終わる
		if d.checkExitWinningProcess(numberOfWinningTickets, numberOfWinningTicketCnt) {
			break
		}
		// 当選したチケットのインデックスをランダムに出力
		index := random.RandomInt(numberOfWinningTicketCnt)

		userBookTicket := userBookTickets[index]
		// 残りの枚数に対して、購入枚数が足りない場合はスルーする
		if numberOfWinningTicketCnt <= userBookTicket.NumberOfTickets {
			continue
		}

		numberOfTickets, err := d.winningProcess(db /* userBookTicket,*/, userBookTicketsTemp[userBookTicket.ID])
		if err != nil {
			// 現状は特に理由なしでエラーになれば当選予定枚数を減らす
			numberOfWinningTicketCnt--
			continue
		}

		winningUserBookTickets[userBookTicket.ID] = true
		// 当選したチケット枚数を管理する
		numberOfWinningTickets += numberOfTickets

		numberOfWinningTicketCnt = d.recalculationWinningTicketCnt(numberOfWinningTicketCnt, numberOfWinningTickets, userBookTickets, winningUserBookTickets)
	}
	// 落選予定枚数 = 申し込み枚数 - 当選結果枚数
	numberOfNotWonTicketCnt = numberOfAppliedTicketCnt - numberOfWinningTickets
	// 落選チケット
	numberOfNotWonTickets := 0

	// 落選処理
	for _, userBookTicket := range userBookTickets {

		// 当選したチケットはtrueのもの
		if winningUserBookTickets[userBookTicket.ID] {
			continue
		}

		for _, userHasTicket := range userBookTicketsTemp[userBookTicket.ID] {
			userHasTicket.TicketStatus = d.ticketStatus.NotWon
			_, err := d.userHasTicket.Save(db, userHasTicket)
			if err != nil {
				continue
			}
			numberOfNotWonTickets++
		}
	}

	if numberOfNotWonTicketCnt != numberOfNotWonTickets {
		log.Println("落選枚数が一致しない")
	}

	return nil
}

func (d *drawInteractor) checkExitWinningProcess(numberOfWinningTicketCnt, numberOfWinningTickets int) bool {
	return numberOfWinningTicketCnt == numberOfWinningTickets
}

// func (d *drawInteractor) checkExitRejectionProcess() bool {
// 	return true
// }

func (d *drawInteractor) winningProcess(db bun.IDB /* userBookTicket *models.UserBookTickets, */, userHasTickets []*models.UserHasTickets) (int, error) {
	numberOfTickets := 0 // 当選処理をした枚数
	// データの整合性はとれているものと仮定する
	// 今は当選した処理を記述
	for _, userHasTicket := range userHasTickets {
		userHasTicket.TicketStatus = d.ticketStatus.Winning
		if _, err := d.userHasTicket.Save(db, userHasTicket); err != nil {
			continue
		}
		numberOfTickets++
	}

	return numberOfTickets, nil
}

// func (d *drawInteractor) notWonProcess(db bun.IDB /* userBookTicket *models.UserBookTickets,*/, userHasTickets []*models.UserHasTickets) (int, error) {
// 	numberOfTickets := 0 // 落選処理をした枚数
// 	// 今は落選した処理を記述
// 	for _, userHasTicket := range userHasTickets {
// 		userHasTicket.TicketStatus = d.ticketStatus.NotWon
// 		if _, err := d.userHasTicket.Save(db, userHasTicket); err != nil {
// 			continue
// 		}
// 		numberOfTickets++
// 	}

// 	return numberOfTickets, nil
// }

func (d *drawInteractor) recalculationWinningTicketCnt(
	numberOfWinningTicketCnt, numberOfWinningTickets int,
	userBookTickets []*models.UserBookTickets,
	winningUserBookTickets map[int]bool,
) int {
	//　残り = 設定当選枚数 - 当選枚数
	cnt := numberOfWinningTicketCnt - numberOfWinningTickets

	isFound := false

	for _, userBookTicket := range userBookTickets {
		// 当選していないもの
		if !winningUserBookTickets[userBookTicket.ID] {
			if userBookTicket.NumberOfTickets <= cnt {
				isFound = true
			}
		}
		// 残りより少ない枚数のチケットがあれば、for文を抜ける
		if isFound {
			break
		}
	}
	// 残りの枚数より少ないものがなければ現状の当選枚数を返す
	if !isFound {
		return numberOfWinningTickets
	}
	// 残りの枚数より少ないものがあれば現状の設定の当選枚数を返す
	return numberOfWinningTicketCnt
}

// func (d *drawInteractor) validationWinningTicket(db bun.IDB, winningTicket *models.UserHasTickets) error {
// 	// 各種データの確認

// 	// 複数枚購入の場合、どちらも当選または落選させたいのでその処理も記述する <-これは元の関数で行う？

// 	return nil
// }
