package models

type Customer struct {
	CustomerNumber         uint64  `json:"customer_number"`
	CustomerName           string  `json:"customer_name"`
	ContactLastName        string  `json:"contact_last_name"`
	ContactFirstName       string  `json:"contact_first_name"`
	Phone                  string  `json:"phone"`
	AddressLine1           string  `json:"address_line_1"`
	AddressLine2           string  `json:"address_line_2"`
	City                   string  `json:"city"`
	State                  string  `json:"state"`
	PostalCode             string  `json:"postal_code"`
	Country                string  `json:"country"`
	SalesRepEmployeeNumber uint64  `json:"sales_rep_employee_number"`
	CreditLimit            float64 `json:"credit_limit"`
}
