package repositories

import (
	"finance-backend/config"
	"finance-backend/models"
)

func CreateRecord(record models.Record) error {
	query := `INSERT INTO records (id, user_id, amount, type, category, notes)
			  VALUES ($1,$2,$3,$4,$5,$6)`

	_, err := config.DB.Exec(query,
		record.ID, record.UserID, record.Amount,
		record.Type, record.Category, record.Notes,
	)

	return err
}

func GetRecords() ([]models.Record, error) {
	rows, err := config.DB.Query("SELECT id, amount, type, category FROM records")
	if err != nil {
		return nil, err
	}

	var records []models.Record

	for rows.Next() {
		var r models.Record
		rows.Scan(&r.ID, &r.Amount, &r.Type, &r.Category)
		records = append(records, r)
	}

	return records, nil
}
