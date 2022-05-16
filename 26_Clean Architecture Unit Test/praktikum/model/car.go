package model

type Car struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Type     string `json:"type"`
	Brand    string `json:"brand"`
	Price    string `json:"price"`
	Rentcar  Rent
	Depotcar Depot
}
