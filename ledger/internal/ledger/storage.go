package ledger

import (
	"errors"
	"sync"
)

type LedgerService struct {
	mu           sync.Mutex
	transactions []Transaction
}

func NewLedger() *LedgerService {
	return &LedgerService{
		transactions: make([]Transaction, 0),
	}
}

func (l *LedgerService) AddTransaction(tx Transaction) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if tx.Amount == 0 {
		return errors.New("transaction amount cannot be zero")
	}

	tx.ID = len(l.transactions) + 1
	l.transactions = append(l.transactions, tx)
	return nil
}

func (l *LedgerService) ListTransactions() []Transaction {
	l.mu.Lock()
	defer l.mu.Unlock()

	transactionsCopy := make([]Transaction, len(l.transactions))
	copy(transactionsCopy, l.transactions)
	return transactionsCopy
}
