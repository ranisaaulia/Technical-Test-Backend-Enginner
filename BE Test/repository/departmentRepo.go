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
	table1 = "Master.Department"
)

func GetAllDataDepartment(ctx context.Context) ([]models.Department, error) {
	var departments []models.Department
	db, err := config.PostgreSQL()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	queryText := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", table1)
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dept models.Department
		var createdAt, updatedAt time.Time
		if err := rows.Scan(&dept.ID, &dept.Name, &createdAt, &dept.CreatedBy, &updatedAt, &dept.UpdatedBy, &dept.DeletedAt); err != nil {
			return nil, err
		}

		dept.CreatedAt = createdAt
		dept.UpdatedAt = updatedAt
		dept.DeletedAt = &time.Time{}

		departments = append(departments, dept)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

func InsertDepartment(ctx context.Context, department models.Department) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("INSERT INTO %v (Department_name, Created_at, Created_by, Updated_at, Updated_by) VALUES ($1, NOW(), $2, NOW(), $3)", table1)
	_, err = db.ExecContext(ctx, queryText, department.Name, department.CreatedBy, department.UpdatedBy)

	if err != nil {
		return err
	}
	return nil
}

func UpdateDepartment(ctx context.Context, department models.Department, id string) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("UPDATE %v SET Department_name = $1, Updated_at = NOW(), Updated_by = $2 WHERE Department_id = $3",
		table1)

	_, err = db.ExecContext(ctx, queryText, department.Name, department.UpdatedBy, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteDepartment(ctx context.Context, id string) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("DELETE FROM %v WHERE Department_id = $1", table1)

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
