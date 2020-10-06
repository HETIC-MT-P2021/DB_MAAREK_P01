package models

type Employee struct {
	EmployeeNumber string `json:"employee_number"`
	LastName       string `json:"last_name"`
	FirstName      string `json:"first_name"`
	Extension      string `json:"extension"`
	Email          string `json:"email"`
	OfficeCode     string `json:"office_code,omitempty"`
	ReportsTo      string `json:"reports_to,omitempty"`
	JobTitle       string `json:"jobs_title"`
}
