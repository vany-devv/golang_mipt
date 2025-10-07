package main

import (
	"fmt"
	"ledger/internal/ledger"
	"time"
)

func main() {
	fmt.Println("Ledger service started")

	service := ledger.NewLedger()

	_ = service.AddTransaction(ledger.Transaction{
		Amount:      150.75,
		Category:    "Food",
		Description: "Lunch at cafe",
		Date:        time.Now(),
	})

	_ = service.AddTransaction(ledger.Transaction{
		Amount:      3000,
		Category:    "Rent",
		Description: "Monthly apartment rent",
		Date:        time.Now(),
	})

	_ = service.AddTransaction(ledger.Transaction{
		Amount:      50,
		Category:    "Transport",
		Description: "Metro tickets",
		Date:        time.Now(),
	})

	for _, tx := range service.ListTransactions() {
		fmt.Printf(
			"ID: %d | Amount: %.2f | Category: %s | Description: %s | Date: %s\n",
			tx.ID, tx.Amount, tx.Category, tx.Description, tx.Date.Format("2006-01-02 15:04"),
		)
	}
}
