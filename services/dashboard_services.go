package services

import "finance-backend/config"

func GetSummary() (float64, float64) {
	var income, expense float64

	config.DB.QueryRow("SELECT COALESCE(SUM(amount),0) FROM records WHERE type='income'").Scan(&income)
	config.DB.QueryRow("SELECT COALESCE(SUM(amount),0) FROM records WHERE type='expense'").Scan(&expense)

	return income, expense
}
