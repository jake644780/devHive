package models

type Idea struct{
	Id int `json:"id"`
	Category string `json:"category"`
	Content string `json:"content"`
}