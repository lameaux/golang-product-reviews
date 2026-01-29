package model

type Review struct {
	ID        ID
	ProductID ID
	FirstName string
	LastName  string
	Review    string
	Rating    Rating
}
