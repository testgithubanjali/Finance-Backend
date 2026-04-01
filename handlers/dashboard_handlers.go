package handlers

import (
	"finance-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSummary(c *gin.Context) {
	income, expense := services.GetSummary()

	c.JSON(http.StatusOK, gin.H{
		"income":  income,
		"expense": expense,
		"net":     income - expense,
	})
}
