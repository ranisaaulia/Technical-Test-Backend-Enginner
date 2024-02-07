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
	table3 = "Master.Location"
)

func GetAllDataLocation(ctx context.Context) ([]models.Location, error) {
	var locations []models.Location
	db, err := config.PostgreSQL()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	queryText := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", table3)
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var loc models.Location
		var createdAt, updatedAt time.Time
		if err := rows.Scan(&loc.ID, &loc.Name, &createdAt, &loc.CreatedBy, &updatedAt, &loc.UpdatedBy, &loc.DeletedAt); err != nil {
			return nil, err
		}

		loc.CreatedAt = createdAt
		loc.UpdatedAt = updatedAt
		loc.DeletedAt = &time.Time{}

		locations = append(locations, loc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}

func InsertLocation(ctx context.Context, location models.Location) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("INSERT INTO %v (Location_Name, Created_at, Created_by, Updated_at, Updated_by) VALUES ($1, NOW(), $2, NOW(), $3)", table3)
	_, err = db.ExecContext(ctx, queryText, location.Name, location.CreatedBy, location.UpdatedBy)

	if err != nil {
		return err
	}
	return nil
}

func UpdateLocation(ctx context.Context, location models.Location, id string) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("UPDATE %v SET Location_Name = $1, Updated_at = NOW(), Updated_by = $2 WHERE Location_id = $3", table3)

	_, err = db.ExecContext(ctx, queryText, location.Name, location.UpdatedBy, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteLocation(ctx context.Context, id string) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("DELETE FROM %v WHERE Location_id = $1", table3)

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
