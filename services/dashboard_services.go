package services

import (
	"finance-backend/config"
	"log"
)

func GetSummary() (float64, float64, error) {
	log.Println("DashboardService: GetSummary started")
	var income, expense float64

	if err := config.DB.QueryRow("SELECT COALESCE(SUM(amount),0) FROM records WHERE type='income'").Scan(&income); err != nil {
		log.Printf("DashboardService: income query error: %v", err)
		return 0, 0, err
	}

	if err := config.DB.QueryRow("SELECT COALESCE(SUM(amount),0) FROM records WHERE type='expense'").Scan(&expense); err != nil {
		log.Printf("DashboardService: expense query error: %v", err)
		return 0, 0, err
	}

	log.Printf("DashboardService: GetSummary completed income=%.2f expense=%.2f", income, expense)
	return income, expense, nil
}
