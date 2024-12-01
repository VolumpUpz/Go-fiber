package main

import (
	"github.com/gofiber/fiber/v2"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DepartmentEmployeeRequest struct {
	DepartmentDetail Department `json:"department"`
	EmployeeDetail   Employee   `json:"employee"`
}

func getDeptEmp(c *fiber.Ctx) error {
	return c.JSON(deptEmp)
}
