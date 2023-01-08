package model

import "time"

type (
	SourceProduct struct {
		ID           uint64    `json:"id"`
		ProductName  string    `json:"product_name"`
		Qty          int       `json:"qty"`
		SellingPrice float64   `json:"selling_price"`
		PromoPrice   float64   `json:"promo_price"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
