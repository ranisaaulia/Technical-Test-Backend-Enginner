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
	table5 = "Attendance"
)

func GetAllAttendance(ctx context.Context) ([]models.Attendance, error) {
	var attendances []models.Attendance
	db, err := config.PostgreSQL()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	queryText := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", table5)
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var att models.Attendance
		var createdAt, updatedAt time.Time
		if err := rows.Scan(&att.ID, &att.EmployeeID, &att.LocationID, &att.AbsentIn, &att.AbsentOut, &createdAt, &att.CreatedBy, &updatedAt, &att.UpdatedBy, &att.DeletedAt); err != nil {
			return nil, err
		}

		att.CreatedAt = createdAt
		att.UpdatedAt = updatedAt
		att.DeletedAt = &time.Time{}

		attendances = append(attendances, att)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}

func InsertAttendance(ctx context.Context, attendance models.Attendance) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("INSERT INTO %v (Employee_id, Location_id, Absent_in, Absent_out, Created_at, Created_by, Updated_at, Updated_by) VALUES ($1, $2, $3, $4, NOW(), $5, NOW(), $6)", table5)
	_, err = db.ExecContext(ctx, queryText, attendance.EmployeeID, attendance.LocationID, attendance.AbsentIn, attendance.AbsentOut, attendance.CreatedBy, attendance.UpdatedBy)

	if err != nil {
		return err
	}
	return nil
}

func UpdateAttendance(ctx context.Context, attendance models.Attendance, id string) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("UPDATE %v SET Employee_id = $1, Location_id = $2, Absent_in = $3, Absent_out = $4, Updated_at = NOW(), Updated_by = $5 WHERE Attendance_id = $6", table5)

	_, err = db.ExecContext(ctx, queryText, attendance.EmployeeID, attendance.LocationID, attendance.AbsentIn, attendance.AbsentOut, attendance.UpdatedBy, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAttendance(ctx context.Context, id string) error {
	db, err := config.PostgreSQL()
	if err != nil {
		log.Fatal("Tidak dapat terhubung ke database", err)
	}
	defer db.Close()

	queryText := fmt.Sprintf("DELETE FROM %v WHERE Attendance_id = $1", table5)

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

func GetAttendanceReport(startDate, endDate string) ([]models.AttendanceReport, error) {
	var attendanceReports []models.AttendanceReport
	db, err := config.PostgreSQL()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	queryText := `
        SELECT
            A.Absent_in,
            A.Absent_out,
            E.Employee_code,
            E.Employee_name,
            D.Department_name,
            P.Position_name,
            L.Location_name
        FROM
            Attendance A
        JOIN
            Employee E ON A.Employee_id = E.Employee_id
        JOIN
            Master.Department D ON E.Department_id = D.Department_id
        JOIN
            Master.Position P ON E.Position_id = P.Position_id
        JOIN
            Master.Location L ON A.Location_id = L.Location_id
        ORDER BY
            A.Absent_in;
    `
	rows, err := db.Query(queryText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attendanceReport models.AttendanceReport
		if err := rows.Scan(
			&attendanceReport.AbsentIn,
			&attendanceReport.AbsentOut,
			&attendanceReport.EmployeeCode,
			&attendanceReport.EmployeeName,
			&attendanceReport.DepartmentName,
			&attendanceReport.PositionName,
			&attendanceReport.LocationName,
		); err != nil {
			return nil, err
		}

		attendanceReports = append(attendanceReports, attendanceReport)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attendanceReports, nil
}
