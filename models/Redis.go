package models

type  User struct { // Redis Model
   Id int64 `json:"id"`
   IIN string `json:"iin"`
   Name string `json:"name"`
   Phone string `json:"phone"`
   Age int64 `json:"age"`
}
