package productmanager

import (
	"github.com/lameaux/golang-product-reviews/database"
	"github.com/lameaux/golang-product-reviews/dto"
	"github.com/lameaux/golang-product-reviews/model"
)

var _ Manager = (*DAOManager)(nil)

type DAOManager struct {
	dao database.DAO
}

func New(dao database.DAO) *DAOManager {
	return &DAOManager{dao: dao}
}

func (*DAOManager) CreateProduct(p *dto.Product) (model.ID, error) {
	return 0, nil
}

func (*DAOManager) UpdateProduct(productID model.ID, p *dto.Product) error {
	return nil
}

func (*DAOManager) DeleteProduct(productID model.ID) error {
	return nil
}

func (*DAOManager) GetProduct(productID model.ID) *dto.ProductWithRating {
	return nil
}

func (*DAOManager) ListProducts(offset int, limit int) []*dto.ProductWithRating {
	return []*dto.ProductWithRating{}
}

func (*DAOManager) CreateProductReview(productID model.ID, r *dto.Review) (model.ID, error) {
	return 0, nil
}

func (*DAOManager) DeleteProductReview(productID model.ID, reviewID model.ID) error {
	return nil
}

func (*DAOManager) UpdateProductReview(productID model.ID, reviewID model.ID, review *dto.Review) error {
	return nil
}

func (*DAOManager) GetProductReview(productID model.ID, reviewID model.ID) *dto.Review {
	return nil
}

func (*DAOManager) ListProductReviews(productID model.ID, offset int, limit int) []*dto.Review {
	return []*dto.Review{}
}
