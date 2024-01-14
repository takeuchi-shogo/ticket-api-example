package enums

type TicketSaleType int

const (
	FirstCome TicketSaleType = iota
	Lot
)

func (t TicketSaleType) String() string {
	switch t {
	case FirstCome:
		return "firstCome"
	case Lot:
		return "lot"
	default:
		return "Unknown"
	}
}
