package enums

type HowToPay int

const (
	Credit          HowToPay = iota //　クレカ決済
	Convenience                     // コンビニ決済
	Carrier                         // キャリア決済
	BankDebit                       //　銀行口座引き落とし
	BankRedirect                    //　銀行へリダイレクト
	BankTransfer                    //　銀行振込
	BuyNowPayLater                  //　後払い
	RealTimePayment                 //　リアルタイム決済
	Vouchers                        //　店舗支払い
	Wallet                          // ウォレット
)
