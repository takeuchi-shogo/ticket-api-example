package enums

type EventCategory int

const (
	// Concert
	JapaneseMusic EventCategory = iota
	JPop
	KPop
	Idol
	Fes
	// Sports
	BaseBool
	Soccer
	// Event
	// Stage
	// Art
)

func (e EventCategory) String() string {
	switch e {
	case JapaneseMusic:
		return "japanesemusic"
	case JPop:
		return "jpop"
	case KPop:
		return "kpop"
	case Idol:
		return "idol"
	case Fes:
		return "fes"
	case BaseBool:
		return "basebool"
	case Soccer:
		return "soccer"
	default:
		return "Unknown"
	}
}
