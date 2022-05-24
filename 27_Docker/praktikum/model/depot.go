package model

type Depot struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	CarID   int    `json:"carId"`
}
