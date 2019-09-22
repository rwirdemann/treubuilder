package domain

// A Booking represents a single booking.
type Booking struct {
	Amount float64
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
	Accounts map[int]Account
}
