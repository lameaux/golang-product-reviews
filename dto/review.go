package dto

import "github.com/lameaux/golang-product-reviews/model"

type Review struct {
	ID        model.ID     `json:"id"`
	FirstName string       `json:"first_name" validate:"required"`
	LastName  string       `json:"last_name" validate:"required"`
	Review    string       `json:"review" validate:"required"`
	Rating    model.Rating `json:"rating" validate:"required,gte=1,lte=5"`
}
