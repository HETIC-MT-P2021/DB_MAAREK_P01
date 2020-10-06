package repo

import (
	"github.com/JackMaarek/DS/conf"
	"github.com/JackMaarek/DS/models"
	"github.com/JackMaarek/DS/util"
)

func QueryOfficeByID(id uint64) (*models.Office, error) {
	var (
		officeCode   string
		city         string
		phone        string
		addressLine2 util.NullString
		addressLine1 string
		state        util.NullString
		country      string
		postalCode   string
		territory    string
	)
	row := conf.DB.QueryRow("SELECT o.officeCode, o.city, o.phone, o.addressLine1, o.addressLine2, o.state, o.country, o.postalCode, o.territory FROM offices o WHERE o.officeCode=(?)", id)

	err := row.Scan(
		&officeCode,
		&city,
		&phone,
		&addressLine1,
		&addressLine2,
		&country,
		&state,
		&postalCode,
		&territory,
	)

	if err != nil {
		return nil, err
	}

	var employees []*models.Employee

	office := models.Office{
		OfficeCode:   officeCode,
		City:         city,
		Phone:        phone,
		AddressLine1: addressLine1,
		AddressLine2: addressLine2.String,
		State:        state.String,
		Country:      country,
		PostalCode:   postalCode,
		Territory:    territory,
	}

	if employees, err = QueryEmployeesByOffice(&office); err != nil {
		return &office, nil
	}

	office.Employee = employees

	return &office, nil
}
