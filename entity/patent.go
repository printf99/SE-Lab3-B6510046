package entity

import "gorm.io/gorm"

type patient struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age_eiei"`
	Email     string `json:"email"`
	
}
