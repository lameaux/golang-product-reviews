package productmanager

import "github.com/lameaux/golang-product-reviews/database"

type Manager struct {
	dao database.DAO
}

func New(dao database.DAO) *Manager {
	return &Manager{dao: dao}
}
