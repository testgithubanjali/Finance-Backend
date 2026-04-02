package services

import (
	"finance-backend/models"
	"finance-backend/repositories"
	"log"

	"github.com/google/uuid"
)

func CreateRecord(record models.Record, userID string) error {
	log.Println("RecordService: CreateRecord started")
	record.ID = uuid.New().String()
	record.UserID = userID
	log.Printf("RecordService: CreateRecord user_id=%s type=%s amount=%.2f", userID, record.Type, record.Amount)

	err := repositories.CreateRecord(record)
	if err != nil {
		log.Printf("RecordService: CreateRecord failed user_id=%s id=%s error=%v", userID, record.ID, err)
		return err
	}
	log.Printf("RecordService: CreateRecord succeeded user_id=%s id=%s", userID, record.ID)
	return nil
}
func GetRecords(userID string) ([]models.Record, error) {
	log.Println("RecordService: GetRecords started")
	records, err := repositories.GetRecordsByUser(userID)
	if err != nil {
		log.Printf("RecordService: GetRecords failed: %v", err)
		return nil, err
	}
	log.Printf("RecordService: GetRecords fetched %d records", len(records))
	return records, nil
}
func GetFilteredRecords(userID, rType, category string, limit, offset int) ([]models.Record, error) {

	log.Printf("Service: filter+pagination user=%s type=%s category=%s limit=%d offset=%d",
		userID, rType, category, limit, offset)

	return repositories.GetFilteredRecords(userID, rType, category, limit, offset)
}
func UpdateRecord(id string, record models.Record) error {
	log.Printf("Service: updating record id=%s", id)
	return repositories.UpdateRecord(id, record)
}

func DeleteRecord(id string) error {
	log.Printf("Service: deleting record id=%s", id)
	return repositories.DeleteRecord(id)
}
