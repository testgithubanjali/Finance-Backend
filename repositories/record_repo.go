package repositories

import (
	"finance-backend/config"
	"finance-backend/models"
	"log"
	"strconv"
)

func CreateRecord(record models.Record) error {
	log.Printf("RecordRepo: CreateRecord user_id=%s id=%s amount=%.2f type=%s", record.UserID, record.ID, record.Amount, record.Type)
	query := `INSERT INTO records (id, user_id, amount, type, category, notes)
			  VALUES ($1,$2,$3,$4,$5,$6)`

	_, err := config.DB.Exec(query,
		record.ID, record.UserID, record.Amount,
		record.Type, record.Category, record.Notes,
	)
	if err != nil {
		log.Printf("RecordRepo: CreateRecord failed user_id=%s id=%s error=%v", record.UserID, record.ID, err)
		return err
	}
	log.Printf("RecordRepo: CreateRecord succeeded user_id=%s id=%s", record.UserID, record.ID)
	return nil
}
func GetRecordsByUser(userID string) ([]models.Record, error) {
	rows, err := config.DB.Query(
		"SELECT id, amount, type, category FROM records WHERE user_id=$1",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.Record

	for rows.Next() {
		var r models.Record
		rows.Scan(&r.ID, &r.Amount, &r.Type, &r.Category)
		records = append(records, r)
	}

	return records, nil
}
func GetFilteredRecords(userID, rType, category string, limit, offset int) ([]models.Record, error) {

	query := "SELECT id, amount, type, category FROM records WHERE user_id=$1"
	args := []interface{}{userID}
	i := 2

	if rType != "" {
		query += " AND type=$" + strconv.Itoa(i)
		args = append(args, rType)
		i++
	}

	if category != "" {
		query += " AND category=$" + strconv.Itoa(i)
		args = append(args, category)
		i++
	}

	query += " LIMIT $" + strconv.Itoa(i) + " OFFSET $" + strconv.Itoa(i+1)
	args = append(args, limit, offset)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.Record

	for rows.Next() {
		var r models.Record
		rows.Scan(&r.ID, &r.Amount, &r.Type, &r.Category)
		records = append(records, r)
	}

	return records, nil
}
func UpdateRecord(id string, record models.Record) error {
	query := `
	UPDATE records
	SET amount=$1, type=$2, category=$3, notes=$4
	WHERE id=$5
	`
	_, err := config.DB.Exec(query,
		record.Amount, record.Type, record.Category, record.Notes, id,
	)
	return err
}
func DeleteRecord(id string) error {
	_, err := config.DB.Exec("DELETE FROM records WHERE id=$1", id)
	return err
}
func GetUserByID(id string) (models.User, error) {
	var user models.User

	err := config.DB.QueryRow(
		"SELECT id, email, role, is_active FROM users WHERE id=$1",
		id,
	).Scan(&user.ID, &user.Email, &user.Role, &user.IsActive)

	return user, err
}
