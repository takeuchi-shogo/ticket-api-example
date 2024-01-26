package enums

type PaymentMethods struct {
	Credit          string //　クレカ決済
	Convenience     string // コンビニ決済
	Carrier         string // キャリア決済
	BankDebit       string //　銀行口座引き落とし
	BankRedirect    string //　銀行へリダイレクト
	BankTransfer    string //　銀行振込
	BuyNowPayLater  string //　後払い
	RealTimePayment string //　リアルタイム決済
	Vouchers        string //　店舗支払い
	Wallet          string // ウォレット
}
