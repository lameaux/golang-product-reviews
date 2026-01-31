package model

const TableProducts = "products"

type Product struct {
	ID          ID           `gorm:"primaryKey;column:id"`
	Name        string       `gorm:"column:product_name"`
	Description string       `gorm:"column:description"`
	Price       PriceInCents `gorm:"column:price"`
}

func (Product) TableName() string {
	return TableProducts
}
