package models

type Person struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FullName    string `json:"full_name"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
}
