package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Transaction represents a bank transaction
type Transaction struct {
	ID          int
	Description string
	Amount      float64
	Date        time.Time
}

// BankAccount represents a bank account with transactions
type BankAccount struct {
	ID           int
	AccountOwner string
	Balance      float64
	Transactions []Transaction
}

var accounts []BankAccount

func getAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	account, err := findAccountByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func addTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	account, err := findAccountByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	var transaction struct {
		Description string  `json:"description"`
		Amount      float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update account balance
	account.Balance += transaction.Amount

	// Add transaction to account
	newTransaction := Transaction{
		Description: transaction.Description,
		Amount:      transaction.Amount,
		Date:        time.Now(),
	}
	account.Transactions = append(account.Transactions, newTransaction)

	c.JSON(http.StatusCreated, account)
}

func findAccountByID(id int) (*BankAccount, error) {
	for _, acc := range accounts {
		if acc.ID == id {
			return &acc, nil
		}
	}
	return nil, fmt.Errorf("account not found")
}
