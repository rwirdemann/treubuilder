package domain

import "math/big"

// A Booking represents a single booking.
type Booking struct {
	Amount big.Float
}

// An Account represents an account with bookings.
type Account struct {
	ID       int
	Owner    string
	Bookings []Booking
}

// An AccountSystem represents a group of related Accounts.
type AccountSystem struct {
	Name     string
	Accounts []Account
}
