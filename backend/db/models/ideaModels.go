package models

type Idea struct{
	Id int `json:"id" gorm:"primaryKey;autoIncrement"`
	Category string `json:"category" gorm:"type:varchar(50);not null"`
	Content string `json:"content" gorm:"type:varchar(50);not null"`
}