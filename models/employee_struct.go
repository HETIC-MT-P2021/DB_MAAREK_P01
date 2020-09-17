package models

type Employee struct {
	EmployeeNumber		string		`json:"employee_number"`
	LastName			string		`json:"last_name"`
	FirstName   		string		`json:"first_name"`
	Extension   		string		`json:"extension"`
	Email   			string		`json:"email"`
	Office				Office		`json:"_"`
	ReportsTo			*Employee	`json:"_,omitempty"`
	JobTitle			string		`json:"jobsTitle"`
}
