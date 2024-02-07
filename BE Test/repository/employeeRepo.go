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
	table4 = "Employee"
)

func GetAllEmployee(ctx context.Context) ([]models.Employee, error) {
	var employees []models.Employee
	db, err := config.PostgreSQL()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	queryText := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", table4)
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var emp models.Employee
		var createdAt, updatedAt time.Time
		if err := rows.Scan(&emp.ID, &emp.Code, &emp.Name, &emp.Password, &emp.DepartmentID, &emp.PositionID, &emp.SuperiorID, &createdAt, &emp.CreatedBy, &updatedAt, &emp.UpdatedBy, &emp.DeletedAt); err != nil {
			return nil, err
		}

		emp.CreatedAt = createdAt
		emp.UpdatedAt = updatedAt
		emp.DeletedAt = &time.Time{}

		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func InsertEmployee(ctx context.Context, employee models.Employee) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("INSERT INTO %v (Employee_code, Employee_name, Password, Department_id, Position_id, Superior, Created_at, Created_by, Updated_at, Updated_by) VALUES ($1, $2, $3, $4, $5, $6, NOW(), $7, NOW(), $8)", table4)
	_, err = db.ExecContext(ctx, queryText, employee.Code, employee.Name, employee.Password, employee.DepartmentID, employee.PositionID, employee.SuperiorID, employee.CreatedBy, employee.UpdatedBy)

	if err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(ctx context.Context, employee models.Employee, id string) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("UPDATE %v SET Employee_code = $1, Employee_name = $2, Password = $3, Department_id = $4, Position_id = $5, Superior = $6, Updated_at = NOW(), Updated_by = $7 WHERE Employee_id = $8", table4)

	_, err = db.ExecContext(ctx, queryText, employee.Code, employee.Name, employee.Password, employee.DepartmentID, employee.PositionID, employee.SuperiorID, employee.UpdatedBy, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEmployee(ctx context.Context, id string) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("DELETE FROM %v WHERE Employee_id = $1", table4)

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
