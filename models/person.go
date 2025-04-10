package models

type Person struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FullName    string    `json:"full_name"`
	Gender      string    `json:"gender"`
	Nationality string    `json:"nationality"`
	Age         int       `json:"age"`
	Friends     []*Person `gorm:"many2many:person_friends;" json:"friends"`
}

// CreatePersonInput структура запроса на создание человека
type CreatePersonInput struct {
	FullName string `json:"full_name" example:"Ivan Ivanov"`
}

// FriendInput структура для запроса добавления друга
type FriendInput struct {
	FriendID uint `json:"friend_id" example:"2"`
}
