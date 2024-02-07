CREATE SCHEMA Master;

CREATE TABLE Master.Department (
    Department_id SERIAL PRIMARY KEY,
    Department_name VARCHAR(255),
    Created_at TIMESTAMP,
    Created_by VARCHAR(255),
    Updated_at TIMESTAMP,
    Updated_by VARCHAR(255),
    Deleted_at TIMESTAMP
);

INSERT INTO Master.Department (Department_name, Created_at, Created_by, Updated_at, Updated_by)
VALUES 
    ('Departemen A', '2024-02-06 08:00:00', 'Admin', '2024-02-06 08:00:00', 'Admin'),
    ('Departemen B', '2024-02-06 08:10:00', 'Admin', '2024-02-06 08:10:00', 'Admin'),
    ('Departemen C', '2024-02-06 08:20:00', 'Admin', '2024-02-06 08:20:00', 'Admin');


CREATE TABLE Master.Position (
    Position_id SERIAL PRIMARY KEY,
    Department_id INT,
    Position_name VARCHAR(255),
    Created_at TIMESTAMP,
    Created_by VARCHAR(255),
    Updated_at TIMESTAMP,
    Updated_by VARCHAR(255),
    Deleted_at TIMESTAMP,
    FOREIGN KEY (Department_id) REFERENCES Master.Department(Department_id)
);

INSERT INTO Master.Position (Department_id, Position_name, Created_at, Created_by, Updated_at, Updated_by)
VALUES 
    (1, 'Manager', '2024-02-06 08:00:00', 'Admin', '2024-02-06 08:00:00', 'Admin'),
    (2, 'Supervisor', '2024-02-06 08:10:00', 'Admin', '2024-02-06 08:10:00', 'Admin'),
    (3, 'Staff', '2024-02-06 08:20:00', 'Admin', '2024-02-06 08:20:00', 'Admin');
    

CREATE TABLE Master.Location (
    location_id SERIAL PRIMARY KEY,
    Location_name VARCHAR(255),
    Created_at TIMESTAMP,
    Created_by VARCHAR(255),
    Updated_at TIMESTAMP,
    Updated_by VARCHAR(255),
    Deleted_at TIMESTAMP
);

INSERT INTO Master.Location (Location_name, Created_at, Created_by, Updated_at, Updated_by)
VALUES 
    ('Location A', '2024-02-06 08:00:00', 'Admin', '2024-02-06 08:00:00', 'Admin'),
    ('Location B', '2024-02-06 08:10:00', 'Admin', '2024-02-06 08:10:00', 'Admin'),
    ('Location C', '2024-02-06 08:20:00', 'Admin', '2024-02-06 08:20:00', 'Admin');


CREATE TABLE Employee (
    Employee_id SERIAL PRIMARY KEY,
    Employee_code VARCHAR(10),
    Employee_name VARCHAR(255),
    Password VARCHAR(255),
    Department_id INT,
    Position_id INT,
    Superior INT,
    Created_at TIMESTAMP,
    Created_by VARCHAR(255),
    Updated_at TIMESTAMP,
    Updated_by VARCHAR(255),
    Deleted_at TIMESTAMP,
    FOREIGN KEY (Department_id) REFERENCES Master.Department(Department_id),
    FOREIGN KEY (Position_id) REFERENCES Master.Position(Position_id),
    FOREIGN KEY (Superior) REFERENCES Employee(Employee_id)
);

INSERT INTO Employee (Employee_code, Employee_name, Password, Department_id, Position_id, Superior, Created_at, Created_by, Updated_at, Updated_by)
VALUES 
    ('22010001', 'Employee A', 'password1', 1, 1, 2, '2024-02-06 08:00:00', 'Admin', '2024-02-06 08:00:00', 'Admin'),
    ('22010002', 'Employee B', 'password2', 2, 2, 1, '2024-02-06 08:10:00', 'Admin', '2024-02-06 08:10:00', 'Admin'),
    ('22010003', 'Employee C', 'password3', 3, 3, 1, '2024-02-06 08:20:00', 'Admin', '2024-02-06 08:20:00', 'Admin');


CREATE TABLE Attendance (
    Attendance_id SERIAL PRIMARY KEY,
    Employee_id INT,
    Location_id INT,
    Absent_in TIMESTAMP,
    Absent_out TIMESTAMP,
    Created_at TIMESTAMP,
    Created_by VARCHAR(255),
    Updated_at TIMESTAMP,
    Updated_by VARCHAR(255),
    Deleted_at TIMESTAMP,
    FOREIGN KEY (Employee_id) REFERENCES Employee(Employee_id),
    FOREIGN KEY (Location_id) REFERENCES Master.Location(location_id)
);

INSERT INTO Attendance (Employee_id, Location_id, Absent_in, Absent_out, Created_at, Created_by, Updated_at, Updated_by)
VALUES 
    (1, 1, '2024-02-06 08:00:00', '2024-02-06 17:00:00', '2024-02-06 08:00:00', 'Admin', '2024-02-06 08:00:00', 'Admin'),
    (2, 2, '2024-02-06 08:30:00', '2024-02-06 17:30:00', '2024-02-06 08:30:00', 'Admin', '2024-02-06 08:30:00', 'Admin'),
    (3, 3, '2024-02-06 09:00:00', '2024-02-06 18:00:00', '2024-02-06 09:00:00', 'Admin', '2024-02-06 09:00:00', 'Admin');


SELECT *FROM ATTENDANCE;

SELECT *FROM EMPLOYEE;

SELECT *FROM MASTER.DEPARTMENT;

SELECT *FROM MASTER.LOCATION;

SELECT *FROM MASTER.POSITION;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


