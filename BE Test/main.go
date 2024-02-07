package main

import (
	"beTest/models"
	"beTest/repository"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// DEPARTMENT
	app.Get("/department", GetDepartment)

	app.Post("/department", AddDepartment)

	app.Put("/department/:id", UpdateDepartment)

	app.Delete("/department/:id/delete", DeleteDepartment)

	// POSITION
	app.Get("/position", GetPosition)

	app.Post("/position", AddPosition)

	app.Put("/position/:id", UpdatePosition)

	app.Delete("/position/:id/delete", DeletePosition)

	// LOCATION

	app.Get("/location", GetLocation)

	app.Post("/location", AddLocation)

	app.Put("/location/:id", UpdateLocation)

	app.Delete("/location/:id/delete", DeleteLocation)

	// EMPLOYEE

	app.Get("/employee", GetEmployee)

	app.Post("/employee", AddEmployee)

	app.Put("/employee/:id", UpdateEmployee)

	app.Delete("/employee/:id/delete", DeleteEmployee)

	// ATTENDANCE

	app.Get("/attendance", GetAttendance)

	app.Get("/attendance-report", GetAttendanceReport)

	app.Post("/attendance", AddAttendance)

	app.Put("/attendance/:id", UpdateAttendance)

	app.Delete("/attendance/:id/delete", DeleteAttendance)

	fmt.Println("Server Running at Port 8181")
	log.Fatal(app.Listen(":8181"))
}

// DEPARTMENT

func GetDepartment(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	department, err := repository.GetAllDataDepartment(ctx)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(department)
}

func AddDepartment(c *fiber.Ctx) error {
	var dep models.Department

	if err := c.BodyParser(&dep); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := repository.InsertDepartment(ctx, dep); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	res := fiber.Map{"status": "Successfully"}
	return c.Status(http.StatusCreated).JSON(res)
}

func UpdateDepartment(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return c.Status(fiber.StatusBadRequest).SendString("Gunakan content type application / json")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var dep models.Department
	if err := c.BodyParser(&dep); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	idDepartment := c.Params("id")

	if err := repository.UpdateDepartment(ctx, dep, idDepartment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "Successfully",
	})
}

func DeleteDepartment(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	idDepartment := c.Params("id")

	if err := repository.DeleteDepartment(ctx, idDepartment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Successfully",
	})
}

// POSITION

func GetPosition(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	position, err := repository.GetAllDataPosition(ctx)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(position)
}

func AddPosition(c *fiber.Ctx) error {
	var pos models.Position

	if err := c.BodyParser(&pos); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := repository.InsertPosition(ctx, pos); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	res := fiber.Map{"status": "Successfully"}
	return c.Status(http.StatusCreated).JSON(res)
}

func UpdatePosition(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return c.Status(fiber.StatusBadRequest).SendString("Gunakan content type application / json")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var pos models.Position
	if err := c.BodyParser(&pos); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	idPosition := c.Params("id")

	if err := repository.UpdatePosition(ctx, pos, idPosition); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "Successfully",
	})
}

func DeletePosition(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	idPosition := c.Params("id")

	if err := repository.DeletePosition(ctx, idPosition); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Successfully",
	})
}

// LOCATION

func GetLocation(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	location, err := repository.GetAllDataLocation(ctx)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(location)
}

func AddLocation(c *fiber.Ctx) error {
	var loc models.Location

	if err := c.BodyParser(&loc); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := repository.InsertLocation(ctx, loc); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	res := fiber.Map{"status": "Successfully"}
	return c.Status(http.StatusCreated).JSON(res)
}

func UpdateLocation(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return c.Status(fiber.StatusBadRequest).SendString("Gunakan content type application / json")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var loc models.Location
	if err := c.BodyParser(&loc); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	idLocation := c.Params("id")

	if err := repository.UpdateLocation(ctx, loc, idLocation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "Successfully",
	})
}

func DeleteLocation(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	idLocation := c.Params("id")

	if err := repository.DeleteLocation(ctx, idLocation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Successfully",
	})
}

// EMPLOYEE

func GetEmployee(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	employee, err := repository.GetAllEmployee(ctx)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(employee)
}

func AddEmployee(c *fiber.Ctx) error {
	var emp models.Employee

	if err := c.BodyParser(&emp); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := repository.InsertEmployee(ctx, emp); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	res := fiber.Map{"status": "Successfully"}
	return c.Status(http.StatusCreated).JSON(res)
}

func UpdateEmployee(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return c.Status(fiber.StatusBadRequest).SendString("Gunakan content type application / json")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var emp models.Employee
	if err := c.BodyParser(&emp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	idEmployee := c.Params("id")

	if err := repository.UpdateEmployee(ctx, emp, idEmployee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "Successfully",
	})
}

func DeleteEmployee(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	idEmployee := c.Params("id")

	if err := repository.DeleteEmployee(ctx, idEmployee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Successfully",
	})
}

// ATTENDANCE

func GetAttendance(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	attendance, err := repository.GetAllAttendance(ctx)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(attendance)
}

func AddAttendance(c *fiber.Ctx) error {
	var atc models.Attendance

	if err := c.BodyParser(&atc); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := repository.InsertAttendance(ctx, atc); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	res := fiber.Map{"status": "Successfully"}
	return c.Status(http.StatusCreated).JSON(res)
}

func UpdateAttendance(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return c.Status(fiber.StatusBadRequest).SendString("Gunakan content type application / json")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var atc models.Attendance
	if err := c.BodyParser(&atc); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	idAttendance := c.Params("id")

	if err := repository.UpdateAttendance(ctx, atc, idAttendance); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "Successfully",
	})
}

func DeleteAttendance(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	idAttendance := c.Params("id")

	if err := repository.DeleteAttendance(ctx, idAttendance); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Successfully",
	})
}

func GetAttendanceReport(c *fiber.Ctx) error {

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	attendanceReport, err := repository.GetAttendanceReport(startDate, endDate)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(attendanceReport)
}
