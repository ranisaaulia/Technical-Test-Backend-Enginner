package repository

import (
	"beTest/config"
	"beTest/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table2 = "Master.Position"
)

func GetAllDataPosition(ctx context.Context) ([]models.Position, error) {
	var positions []models.Position
	db, err := config.PostgreSQL()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	queryText := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", table2)
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pos models.Position
		var createdAt, updatedAt time.Time
		if err := rows.Scan(&pos.ID, &pos.DepartmentID, &pos.Name, &createdAt, &pos.CreatedBy, &updatedAt, &pos.UpdatedBy, &pos.DeletedAt); err != nil {
			return nil, err
		}

		pos.CreatedAt = createdAt
		pos.UpdatedAt = updatedAt
		pos.DeletedAt = &time.Time{}

		positions = append(positions, pos)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}

func InsertPosition(ctx context.Context, position models.Position) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("INSERT INTO %v (Department_id, Position_name, Created_at, Created_by, Updated_at, Updated_by) VALUES ($1, $2, NOW(), $3, NOW(), $4)", table2)
	_, err = db.ExecContext(ctx, queryText, position.DepartmentID, position.Name, position.CreatedBy, position.UpdatedBy)

	if err != nil {
		return err
	}
	return nil
}

func UpdatePosition(ctx context.Context, position models.Position, id string) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("UPDATE %v SET Position_name = $1, Updated_at = NOW(), Updated_by = $2 WHERE Position_id = $3",
		table2)

	_, err = db.ExecContext(ctx, queryText, position.Name, position.UpdatedBy, id)
	if err != nil {
		return err
	}

	return nil
}

func DeletePosition(ctx context.Context, id string) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("DELETE FROM %v WHERE Position_id = $1", table2)

	result, err := db.ExecContext(ctx, queryText, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("id tidak ada")
	}

	return nil
}
