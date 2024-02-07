# Technical-Test-Backend-Enginner

This repository contains the answer code from the test given 

## Task List

1. **API CRUD for Master Data Department**: 
    - GET `/department`: Retrieve all departments.
    - POST `/department`: Create a new department.
    - PUT `/department/:id`: Update an existing department.
    - DELETE `/department/:id/delete`: Delete a department.

2. **API CRUD for Master Data Position**:
    - GET `/position`: Retrieve all positions.
    - POST `/position`: Create a new position.
    - PUT `/position/:id`: Update an existing position.
    - DELETE `/position/:id/delete`: Delete a position.

3. **API CRUD for Master Data Location**:
    - GET `/location`: Retrieve all locations.
    - POST `/location`: Create a new location.
    - PUT `/location/:id`: Update an existing location.
    - DELETE `/location/:id/delete`: Delete a location.

4. **API CRUD for Employee**:
    - GET `/employee`: Retrieve all employees.
    - POST `/employee`: Create a new employee.
    - PUT `/employee/:id`: Update an existing employee.
    - DELETE `/employee/:id/delete`: Delete an employee.

5. **API CRUD for Absent**:
    - GET `/attendance`: Retrieve all attendance records.
    - POST `/attendance`: Record new attendance.
    - PUT `/attendance/:id`: Update an attendance record.
    - DELETE `/attendance/:id/delete`: Delete an attendance record.

6. **API for Reporting Absence based on Time Interval**:
    - GET `/attendance-report`: Retrieve attendance reports based on a specified time interval.
      - Response includes:
        - Date
        - Employee Code
        - Employee Name
        - Department Name
        - Position Name
        - Location Name
        - Absent In
        - Absent Out

**Note**: Due to some challenges, the implementation of the login token using JWT authentication has not been completed yet.

## Setup Instructions

1. Clone the repository.
2. Set up the database using the provided schema and initial data backup.
3. Configure the database connection in the application configuration.
4. Run the application.
5. Import the Postman collection for testing the API endpoints.

## Technologies Used

- Golang
- PostgreSQL / SQL Server
- Go Fiber
