package belajargolanggoroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	balance int
}

func (bankAccount *BankAccount) AddBalance(amount int) {
	bankAccount.RWMutex.Lock()
	bankAccount.balance = bankAccount.balance + amount
	bankAccount.RWMutex.Unlock()
}

func (bankAccount *BankAccount) GetBalance() int {
	bankAccount.RWMutex.RLock()
	balance := bankAccount.balance
	bankAccount.RWMutex.RUnlock()
	return balance
}

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("counter: ", x)
}

func TestRWMutex(testing *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance: ", account.GetBalance())
}
