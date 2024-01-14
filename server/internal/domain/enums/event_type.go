package enums

type EventType int

const (
	Concert EventType = iota
	Sports
	Event
	Stage
	Art
)

func (e EventType) String() string {
	switch e {
	case Concert:
		return "concert"
	case Sports:
		return "sports"
	case Event:
		return "event"
	case Stage:
		return "stage"
	case Art:
		return "art"
	default:
		return "Unknown"
	}
}
