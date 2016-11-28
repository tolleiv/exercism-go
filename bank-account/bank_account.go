package account

import (
	"sync"
)

type Account struct {
	sync.RWMutex
	balance int64
	open bool

}

// Open creates a new bank account with the initial depost
func Open(init int64) *Account {
	if init < 0 {
		return nil
	}
	return &Account{balance: init, open: true}
}

// Close will close the account and return the current balance
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.open {
		a.open = false
		return a.balance, true
	} else {
		return 0, false
	}
}

// Balance will returned the account balance
func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()

	if a.open {
		return a.balance, true
	} else {
		return 0, false
	}
}

// Deposit will deposit or withdraw (if negative) the specified amount
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.open && a.balance + amount >= 0 {
		a.balance += amount
		return a.balance, true
	} else {
		return a.balance, false
	}
}
