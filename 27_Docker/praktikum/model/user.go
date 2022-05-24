package model

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NIK      string `json:"nik"`
	//IsAdmin int
	Rents []Rent `json:"rents" gorm:"foreignKey:UserID"`
}
