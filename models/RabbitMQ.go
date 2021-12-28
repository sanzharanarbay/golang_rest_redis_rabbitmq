package models

type Student struct {
	Id int64 `json:"id"`
	Fullname string `json:"fullname"`
	Group string `json:"group"`
	Age int64 `json:"age"`
	GPA float64 `json:"gpa"`
}