package models

import (
	"github.com/jinzhu/gorm"
)

type Ability struct {
	gorm.Model
	AbilityName string
	AbilityType string
}