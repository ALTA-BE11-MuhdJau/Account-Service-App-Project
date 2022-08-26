package entities

import "time"

type TopUp struct {
	ID        int
	UserID    int
	Saldo     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
