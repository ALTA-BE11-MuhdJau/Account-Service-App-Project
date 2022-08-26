package entities

import "time"

type User struct {
	ID        int
	Name      string
	DoB       string // DoB : Date of Birth (Tanggal Lahir)
	Gender    string
	Telp      string
	Password  string
	SisaSaldo int
	CreatedAt time.Time
	UpdatedAt time.Time
}
