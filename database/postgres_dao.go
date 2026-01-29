package database

import "gorm.io/gorm"

type postgresDAO struct {
	db *gorm.DB
}

func NewPostgresDAO(db *gorm.DB) *postgresDAO {
	return &postgresDAO{
		db: db,
	}
}
