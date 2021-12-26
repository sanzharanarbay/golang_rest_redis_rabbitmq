package models

import "time"

type Order struct{
	Id int64 `json:"id"`
	Code string `json:"code"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
