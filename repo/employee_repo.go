package repo

import (
	"github.com/JackMaarek/DS/conf"
	"github.com/JackMaarek/DS/models"
	"github.com/JackMaarek/DS/util"
)

func QueryEmployeeById(id uint64) (*models.Employee, error) {
	var (
		employeeNumber string
		lastName       string
		firstName      string
		extension      string
		email          string
		officeCode     string
		reportsTo      util.NullString
		jobTitle       string
	)

	row := conf.DB.QueryRow("SELECT e.employeeNumber, e.lastName, e.firstName, e.extension, e.email, e.officeCode, e.reportsTo, e.jobTitle FROM employees e WHERE e.employeeNumber = (?)", id)

	err := row.Scan(
		&employeeNumber,
		&lastName,
		&firstName,
		&extension,
		&email,
		&officeCode,
		&reportsTo,
		&jobTitle,
	)

	if err != nil {
		return nil, err
	}

	employee := &models.Employee{
		EmployeeNumber: employeeNumber,
		LastName:       lastName,
		FirstName:      firstName,
		Extension:      extension,
		Email:          email,
		OfficeCode:     officeCode,
		JobTitle:       jobTitle,
	}

	return employee, nil
}

func QueryEmployeesByOffice(office *models.Office) ([]*models.Employee, error) {
	var (
		employeeNumber string
		lastName       string
		firstName      string
		extension      string
		email          string
		reportsTo      util.NullString
		jobTitle       string
	)

	rows, err := conf.DB.Query("SELECT e.employeeNumber, e.lastName, e.firstName, e.extension, e.email, e.reportsTo, e.jobTitle FROM employees e WHERE e.officeCode = (?)", office.OfficeCode)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*models.Employee

	for rows.Next() {
		err = rows.Scan(
			&employeeNumber,
			&lastName,
			&firstName,
			&extension,
			&email,
			&reportsTo,
			&jobTitle,
		)

		employee := &models.Employee{
			EmployeeNumber: employeeNumber,
			LastName:       lastName,
			FirstName:      firstName,
			Extension:      extension,
			Email:          email,
			ReportsTo:      reportsTo.String,
			JobTitle:       jobTitle,
		}

		employees = append(employees, employee)
	}

	if err != nil {
		return nil, err
	}

	return employees, nil
}
