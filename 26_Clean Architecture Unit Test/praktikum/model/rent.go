package model

type Rent struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Start    string `json:"start"`
	Deadline string `json:"deadline"`
	Total    int    `json:"total"`
	UserID   int    `json:"userId"`
	CarID    int    `json:"carId"`
}
