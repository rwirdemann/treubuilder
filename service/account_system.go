package service

import (
	"github.com/rwirdemann/treubuilder/domain"
)

var accountSystems []domain.AccountSystem

func init() {
	as := domain.AccountSystem{Name: "WEG Winterhuder Weg"}
	as.Accounts = make(map[int]domain.Account)
	a1 := domain.Account{ID: 1, Owner: "Wirdemann"}
	a1.Bookings = append(a1.Bookings, domain.Booking{Amount: 99.23})
	as.Accounts[a1.ID] = a1
	a2 := domain.Account{ID: 2, Owner: "Paura"}
	as.Accounts[a2.ID] = a2
	AddAccountSystem(as)
}

func AddAccountSystem(as domain.AccountSystem) {
	accountSystems = append(accountSystems, as)
}

func GetAccountSystem(id int) domain.AccountSystem {
	return accountSystems[0]
}

func GetAccount(id int) domain.Account {
	return accountSystems[0].Accounts[id]
}
