package database

import (
	"context"
	"fmt"

	"github.com/lameaux/golang-product-reviews/model"
	"gorm.io/gorm"
)

var _ DAO = (*postgresDAO)(nil)

type postgresDAO struct {
	db *gorm.DB
}

func NewPostgresDAO(db *gorm.DB) *postgresDAO {
	return &postgresDAO{
		db: db,
	}
}

func (d *postgresDAO) CreateProduct(ctx context.Context, product *model.Product) (model.ID, error) {
	err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(product).Error; err != nil {
			return fmt.Errorf("tx.Create: %w", err)
		}
		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("CreateProduct: %w", err)
	}

	return product.ID, nil
}

func (d *postgresDAO) UpdateProduct(ctx context.Context, product *model.Product) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(product).Error; err != nil {
			return fmt.Errorf("tx.Save: %w", err)
		}

		return nil
	})
}

func (d *postgresDAO) DeleteProduct(ctx context.Context, id model.ID) error {
	product := &model.Product{ID: id}
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(product).Error; err != nil {
			return fmt.Errorf("tx.Delete: %w", err)
		}

		return nil
	})
}

func (d *postgresDAO) GetProduct(ctx context.Context, id model.ID) (*model.Product, error) {
	var product model.Product

	if err := d.db.WithContext(ctx).
		Table(model.TableProducts).
		Where("id = ?", id).
		Take(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("GetProduct: %w", err)
	}

	return &product, nil
}

func (d *postgresDAO) GetProductRating(ctx context.Context, id model.ID) (float32, error) {
	var rating float32
	if err := d.db.WithContext(ctx).
		Table(model.TableReviews).
		Select("avg(rating)").
		Where("product_id = ?", id).
		Having("count(*) > 0").
		Take(&rating).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, fmt.Errorf("GetProductRating: %w", err)
	}

	return rating, nil
}

func (d *postgresDAO) ListProducts(ctx context.Context, offset int, limit int) ([]*model.Product, error) {
	var result []*model.Product

	if err := d.db.WithContext(ctx).
		Table(model.TableProducts).
		Offset(offset).
		Limit(limit).
		Scan(&result).Error; err != nil {
		return nil, fmt.Errorf("ListProducts: %w", err)
	}

	return result, nil
}

func (d *postgresDAO) CreateProductReview(ctx context.Context, review *model.Review) (model.ID, error) {
	err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(review).Error; err != nil {
			return fmt.Errorf("tx.Create: %w", err)
		}
		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("dao.CreateProductReview: %w", err)
	}

	return review.ID, nil
}

func (d *postgresDAO) UpdateProductReview(ctx context.Context, review *model.Review) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(review).Error; err != nil {
			return fmt.Errorf("tx.Save: %w", err)
		}

		return nil
	})
}

func (d *postgresDAO) DeleteProductReview(ctx context.Context, reviewID model.ID) error {
	review := &model.Review{ID: reviewID}
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(review).Error; err != nil {
			return fmt.Errorf("tx.Delete: %w", err)
		}

		return nil
	})
}

func (d *postgresDAO) GetProductReview(ctx context.Context, id model.ID) (*model.Review, error) {
	var review model.Review

	if err := d.db.WithContext(ctx).
		Table(model.TableReviews).
		Where("id = ?", id).
		Take(&review).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("GetProductReview: %w", err)
	}

	return &review, nil
}

func (d *postgresDAO) ListProductReviews(ctx context.Context, productID model.ID, offset int, limit int) ([]*model.Review, error) {
	var result []*model.Review

	if err := d.db.WithContext(ctx).
		Table(model.TableReviews).
		Offset(offset).
		Limit(limit).
		Scan(&result).Error; err != nil {
		return nil, fmt.Errorf("ListProductReviews: %w", err)
	}

	return result, nil
}
