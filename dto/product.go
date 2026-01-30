package dto

import "github.com/lameaux/golang-product-reviews/model"

type Product struct {
	ID          model.ID           `json:"id"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description" validate:"required"`
	Price       model.PriceInCents `json:"price" validate:"required"`
}

type ProductWithRating struct {
	Product
	Rating float32 `json:"rating"`
}
