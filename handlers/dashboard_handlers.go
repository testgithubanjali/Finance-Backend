package handlers

import (
	"finance-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSummary(c *gin.Context) {
	var income, expense float64

	config.DB.QueryRow("SELECT COALESCE(SUM(amount),0) FROM records WHERE type='income'").Scan(&income)
	config.DB.QueryRow("SELECT COALESCE(SUM(amount),0) FROM records WHERE type='expense'").Scan(&expense)

	c.JSON(http.StatusOK, gin.H{
		"income":  income,
		"expense": expense,
		"net":     income - expense,
	})
}
