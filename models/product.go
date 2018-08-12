package models

import "time"

type ProductEntity struct {
	Id           int64     `json:"id"`
	Sku          string    `json:"sku"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	SpecialPrice float64   `json:"special_price"`
	Description  *string   `json:"description"`
	BarCode      *string   `json:"barcode"`
	CreatedAt    time.Time `json:"created_at"`
	UnitWeight   float64   `json:"unit_weigh"`
	ProductType  *string   `json:"product_type"`
}

type ProductImage struct{}

type ProductReivew struct{}
