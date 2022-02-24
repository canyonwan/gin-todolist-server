package model

import (
	"github.com/jinzhu/gorm"
)

type TodoItem struct {
	gorm.Model
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
