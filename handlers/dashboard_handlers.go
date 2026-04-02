package handlers

import (
	"finance-backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSummary(c *gin.Context) {
	log.Println("Dashboard: GetSummary started")

	income, expense, err := services.GetSummary()
	if err != nil {
		log.Printf("Dashboard: GetSummary failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch summary"})
		return
	}

	netTotal := income - expense
	log.Printf("Dashboard: GetSummary result income=%.2f expense=%.2f net=%.2f", income, expense, netTotal)

	c.JSON(http.StatusOK, gin.H{
		"income":  income,
		"expense": expense,
		"net":     netTotal,
	})
}
