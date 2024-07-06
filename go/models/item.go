package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	Profession string `json:"profession"`
	Location   string `json:"location"`
}
