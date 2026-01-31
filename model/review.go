package model

const TableReviews = "reviews"

type Review struct {
	ID        ID     `gorm:"primaryKey;column:id"`
	ProductID ID     `gorm:"column:product_id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Review    string `gorm:"column:review"`
	Rating    Rating `gorm:"column:rating"`
}

func (Review) TableName() string {
	return TableReviews
}
