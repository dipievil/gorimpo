package ports

import "github.com/LXSCA7/gorimpo/internal/core/domain"

type Notifier interface {
	Send(offer domain.Offer) error
	SendText(message string) error
}
