package model

type MessagingMethod int

const (
	WhatsApp MessagingMethod = iota
	Telegram
	Signal
	Messenger
)

func (m MessagingMethod) Int() int {
	return int(m)
}

// func (m MessagingMethod) String() string {
// 	switch m {
// 	case WhatsApp:
// 		return "WhatsApp"
// 	case Telegram:
// 		return "Telegram"
// 	case Signal:
// 		return "Signal"
// 	case Messenger:
// 		return "Messenger"
// 	default:
// 		return "Unknown"
// 	}
// }
