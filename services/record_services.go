package services

import (
	"finance-backend/models"
	"finance-backend/repositories"

	"github.com/google/uuid"
)

func CreateRecord(record models.Record, userID string) error {
	record.ID = uuid.New().String()
	record.UserID = userID

	return repositories.CreateRecord(record)
}

func GetRecords() ([]models.Record, error) {
	return repositories.GetRecords()
}
