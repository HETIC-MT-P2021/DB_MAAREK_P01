package repo

import (
	"fmt"
	"github.com/JackMaarek/DS/conf"
	"github.com/JackMaarek/DS/models"
	"github.com/JackMaarek/DS/util"
)

func QueryCustomerById(id uint64) (*models.Customer, error) {
	results, err := conf.DB.Query("SELECT * FROM customers where customerNumber = ?", id)
	if err != nil {
		fmt.Println(err.Error()) // proper error handling instead of panic in your app
		return &models.Customer{}, err
	}
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

	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(
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
			fmt.Println(err.Error()) // proper error handling instead of panic in your app
			return &models.Customer{}, err
		}
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
