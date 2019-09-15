package models

type Subscribe struct {
	BaseModel
	Email string `json:"email" gorm:"varchar(40);unique;not null"`
}
