package account

import (
	"sync"
)

// Account represents a bank account
type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

// Open accepts an initial deposit amount and returns a newly opened *Account.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit, closed: false}
}

// Close closes a bank account if the account is open and returns the remaining
// balance and a bool indicating if the account is open.
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	payout := a.balance
	a.balance = 0
	a.closed = true
	return payout, true
}

// Balance returns the balance of an account if the account is open and a bool
// indicating if the account is open.
func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit accepts an amount and deposits this amount into an account if it's
// a positive number and withdraws the amount from the account if it's a negative
// number. It returns the new account balance and a bool indicating if the
// account is open.
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed || amount+a.balance < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
