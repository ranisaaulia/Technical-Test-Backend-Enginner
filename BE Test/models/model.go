package models

import "time"

type Department struct {
	ID        int
	Name      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt *time.Time
}

type Position struct {
	ID           int
	DepartmentID int
	Name         string
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string
	DeletedAt    *time.Time
}

type Location struct {
	ID        int
	Name      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt *time.Time
}

type Employee struct {
	ID           int
	Code         string
	Name         string
	Password     string
	DepartmentID int
	PositionID   int
	SuperiorID   int
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string
	DeletedAt    *time.Time
}

type Attendance struct {
	ID         int
	EmployeeID int
	LocationID int
	AbsentIn   time.Time
	AbsentOut  time.Time
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
	DeletedAt  *time.Time
}

type AttendanceReport struct {
	Date           time.Time
	EmployeeCode   string
	EmployeeName   string
	DepartmentName string
	PositionName   string
	LocationName   string
	AbsentIn       time.Time
	AbsentOut      time.Time
}
