package enums

type UserHasTicketStatus struct {
	// 申し込み状態
	Applied string
	// 抽選中
	StillChoosing string
	// 当選
	Winning string // 決済待ち状態にする

	// 落選
	NotWon string

	// 購入済み
	Purchased string

	// 支払い期限切れ
	ExpiredPaymentByCreditCard       string
	ExpiredPaymentByConvenienceStore string

	// 返金
	Refunding string
	Refunded  string

	// 譲渡
	Transferring string
	Transferred  string

	// その他エラー
	OtherError string

	// 新しいステータス（例）
	CustomStatus string
}

func NewUserHasTicketStatus() *UserHasTicketStatus {
	return &UserHasTicketStatus{
		Applied:                          "applied",
		StillChoosing:                    "still_choosing",
		Winning:                          "winning",
		NotWon:                           "not_won",
		Purchased:                        "purchased",
		ExpiredPaymentByCreditCard:       "expire_payment_by_credit_card",
		ExpiredPaymentByConvenienceStore: "expire_payment_by_convenience_store",
		Refunding:                        "refunding",
		Refunded:                         "refunded",
		Transferring:                     "transferring",
		Transferred:                      "transferred",
		OtherError:                       "other_error",
		CustomStatus:                     "custom_status",
	}
}
