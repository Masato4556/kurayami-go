package entity

import "github.com/jinzhu/gorm"

// Todo todo model
type Todo struct {
	gorm.Model
	Text   string
	Status string
}
