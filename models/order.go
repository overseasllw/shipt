package models

import "time"

type Order struct {
	Id             int64          `json:"id"`
	ConfirmationId int64          `json:"confirmation_id"`
	Status         string         `json:"status"`
	UserId         int64          `json:"customer_id"`
	BaseSubtotal   float64        `json:"base_subtotal"`
	Subtotal       float64        `json:"subtotal"`
	DiscountAmount *float64       `json:"discount_amount"`
	TaxAmount      *float64       `json:"tax_amount"`
	GrandTotal     float64        `json:"grand_total"`
	ShippingAmount *float64       `json:"shipping_amount"`
	CreatedAt      time.Time      `json:"created_at"`
	OrderItems     []OrderItem    `json:"order_items"`
	Address        []OrderAddress `json:"address"`
}

type OrderItem struct {
	Id             int64     `json:"id"`
	OrderId        int64     `json:"sale_order_id"`
	ProductId      int64     `json:"product_id"`
	Qty            *int64    `json:"qty"`
	Price          float64   `json:"price"`
	Weight         *float64  `json:"weight"`
	BaseSubtotal   float64   `json:"base_subtotal"`
	Subtotal       float64   `json:"subtotal"`
	DiscountAmount *float64  `json:"discount_amount"`
	TaxAmount      *float64  `json:"tax_amount"`
	CreatedAt      time.Time `json:"created_at"`
}
type OrderAddress struct {
	Id          int64   `json:"id"`
	AddressType string  `json:"address_type"`
	Fullname    string  `json:"fullname"`
	Company     *string `json:"company"`
	Address1    string  `json:"address1"`
	Address2    *string `json:"address2"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Zip         string  `json:"zip"`
	Zip4        *string `json:"zip4"`
	Country     string  `json:"country"`
}
