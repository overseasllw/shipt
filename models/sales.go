package models

type CategorySales struct {
	CustomerId      int64   `json:"customer_id"`
	FirstName       string  `json:"customer_first_name"`
	CategoryId      int64   `json:"category_id"`
	CategoryName    string  `json:"category_name"`
	NumberPurchased float64 `json:"number_purchased"`
}
type ProductSaleByTimeRange struct {
	ProductId   int64   `json:"product_id"`
	Date        string  `json:"date"`
	ProductName string  `json:"product_name"`
	QtySold     float64 `json:"quantity"`
}
