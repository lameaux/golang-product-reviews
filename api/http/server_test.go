package http

import (
	"github.com/gorilla/mux"
	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/productmanager"
	"github.com/rs/zerolog/log"
)

func testRouter() *mux.Router {
	server := New(0, &log.Logger, stubProductManager())
	return server.CreateRouter()
}

func stubProductManager() *productmanager.StubManager {
	return &productmanager.StubManager{
		Products: []*dto.ProductWithRating{
			{
				Product: dto.Product{
					ID:          1,
					Name:        "P1",
					Description: "P1 desc",
					Price:       100,
				},
				Rating: 1,
			},
		},
		Reviews: []*dto.Review{
			{
				ID:        1,
				FirstName: "Sergej",
				LastName:  "Sizov",
				Review:    "Perfect",
				Rating:    5,
			},
		},
	}
}
