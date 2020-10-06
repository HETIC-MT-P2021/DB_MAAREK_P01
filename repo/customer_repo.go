package repo

import (
	"github.com/JackMaarek/DS/conf"
	"github.com/JackMaarek/DS/models"
	"github.com/JackMaarek/DS/util"
)

func QueryCustomerById(id uint64) (*models.Customer, error) {
	row := conf.DB.QueryRow("SELECT c.customerNumber, c.customerName, c.contactLastName, c.contactFirstName, c.phone, c.addressLine1, c.addressLine2, c.city, c.state, c.postalCode, c.country, c.salesRepEmployeeNumber, c.creditLimit FROM customers c where customerNumber = ?", id)

	var (
		CustomerNumber         uint64
		CustomerName           string
		ContactLastName        string
		ContactFirstName       string
		Phone                  string
		AddressLine1           string
		AddressLine2           util.NullString
		City                   string
		State                  util.NullString
		PostalCode             util.NullString
		Country                string
		SalesRepEmployeeNumber uint64
		CreditLimit            float64
	)

	err := row.Scan(
		&CustomerNumber,
		&CustomerName,
		&ContactLastName,
		&ContactFirstName,
		&Phone,
		&AddressLine1,
		&AddressLine2,
		&City,
		&State,
		&PostalCode,
		&Country,
		&SalesRepEmployeeNumber,
		&CreditLimit,
	)
	if err != nil {
		return nil, err
	}

	c := models.Customer{
		CustomerNumber:         CustomerNumber,
		CustomerName:           CustomerName,
		ContactLastName:        ContactLastName,
		ContactFirstName:       ContactFirstName,
		Phone:                  Phone,
		AddressLine1:           AddressLine1,
		AddressLine2:           AddressLine2.String,
		City:                   City,
		State:                  State.String,
		PostalCode:             PostalCode.String,
		Country:                Country,
		SalesRepEmployeeNumber: SalesRepEmployeeNumber,
		CreditLimit:            CreditLimit,
	}
	return &c, nil
}
