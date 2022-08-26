package entities

import "time"

type Transfer struct {
	ID         int
	SenderID   int
	ReceiverID int
	Saldo      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
