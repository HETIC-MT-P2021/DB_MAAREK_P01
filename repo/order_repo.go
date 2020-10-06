package repo

import (
	"fmt"
	"github.com/JackMaarek/DS/conf"
	"github.com/JackMaarek/DS/models"
	"time"
)

func QueryOrdersByCustomerId(id uint64) (*[]models.PublicOrder, error) {
	var (
		Number      uint64
		ItemsNumber uint64
		TotalPrice  float64
	)
	var orders []models.PublicOrder

	rows, err := conf.DB.Query("SELECT o.orderNumber, SUM(od.quantityOrdered) AS 'Items number', SUM(od.priceEach * od.quantityOrdered) AS 'Total price' FROM orders o INNER JOIN orderdetails od ON od.orderNumber = o.orderNumber WHERE o.customerNumber = ? GROUP BY o.orderNumber;", id)

	defer rows.Close()
	if err != nil {
		return &[]models.PublicOrder{}, err
	}
	for rows.Next() {
		err = rows.Scan(
			&Number,
			&ItemsNumber,
			&TotalPrice,
		)
		fmt.Println(Number)

		order := models.PublicOrder{
			Number:      Number,
			ItemsNumber: ItemsNumber,
			TotalPrice:  TotalPrice,
		}
		orders = append(orders, order)
	}
	if err != nil {
		return nil, err
	}

	return &orders, err
}

func QueryOrderById(id uint64) (*models.Order, error) {
	var (
		Number            uint64
		Date              time.Time
		CustomerFirstName string
		CustomerLastName  string
		ItemsNumber       uint64
		TotalPrice        float64
	)
	row := conf.DB.QueryRow("SELECT o.orderNumber, o.orderDate , c.contactFirstName, c.contactLastName, SUM(od.quantityOrdered) AS 'Items number', SUM(od.priceEach * od.quantityOrdered) AS 'Total price' FROM orders o INNER JOIN customers c ON c.customerNumber = o.customerNumber INNER JOIN orderdetails od ON od.orderNumber = o.orderNumber WHERE o.orderNumber = ? GROUP BY o.orderNumber", id)

	err := row.Scan(
		&Number,
		&Date,
		&CustomerFirstName,
		&CustomerLastName,
		&ItemsNumber,
		&TotalPrice,
	)
	if err != nil {
		return nil, err
	}

	order := models.Order{
		Number:            Number,
		Date:              Date,
		CustomerFirstName: CustomerFirstName,
		CustomerLastName:  CustomerLastName,
		ItemsNumber:       ItemsNumber,
		TotalPrice:        TotalPrice,
	}

	return &order, err
}

func QueryOrderDetailsByOrderNumber(number uint64) (*[]models.OrderDetails, error) {
	var (
		ProductCode        string
		ProductName        string
		ProductDescription string
		Quantity           uint64
		UnitPrice          float64
		LinePrice          float64
	)
	var orderDetails []models.OrderDetails

	rows, err := conf.DB.Query("SELECT od.productCode,p.productName, p.productDescription, od.quantityOrdered, od.priceEach, (od.quantityOrdered*od.priceEach) FROM orders o INNER JOIN orderdetails od ON od.orderNumber = o.orderNumber INNER JOIN products p ON od.productCode = p.productCode WHERE o.orderNumber = ? ORDER BY od.orderLineNumber", number)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&ProductCode,
			&ProductName,
			&ProductDescription,
			&Quantity,
			&UnitPrice,
			&LinePrice,
		)

		orderDetail := models.OrderDetails{
			ProductCode:        ProductCode,
			ProductName:        ProductName,
			ProductDescription: ProductDescription,
			Quantity:           Quantity,
			UnitPrice:          UnitPrice,
			LinePrice:          LinePrice,
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	if err != nil {
		return nil, err
	}

	return &orderDetails, err
}
