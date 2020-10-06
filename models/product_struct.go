package models

type Product struct {
	ProductCode        string       `json:"product_code"`
	ProductName        string       `json:"product_name"`
	ProductLine        *ProductLine `json:"_"`
	ProductScale       string       `json:"product_scale"`
	ProductVendor      string       `json:"product_vendor"`
	ProductDescription string       `json:"product_description"`
	QuantityInStock    uint         `json:"quantity_in_stock"`
	BuyPrice           float64      `json:"buy_price"`
	Msrp               float64      `json:"msrp,omitempty"`
}

type ProductLine struct {
	ProductLine     string `json:"product_line"`
	TextDescription string `json:"text_description,omitempty"`
	HtmlDescription string `json:"html_description,omitempty"`
	Image           []byte `json:"image,omitempty"`
}
