package models

import (
	"github.com/jinzhu/gorm"
)

type Hero struct {
	gorm.Model
	HeroName      string `gorm:"uniqueIndex"`
	HeroType      string
	HeroAbilities []*Ability `gorm:"column:abilities"`
}

func (h *Hero) CreateUser(db *gorm.DB) error {
	return db.Create(h).Error;
}

func GetAllHeros(db *gorm.DB, heros *[]Hero) error {
	return db.Find(&heros).Error;
}

func (h* Hero) DeleteHero(db *gorm.DB) error {
	return db.Delete(h).Error;
}

func (h* Hero) GetDetail(db *gorm.DB, id int) error {
	return db.Find(h, id).Error;
}

