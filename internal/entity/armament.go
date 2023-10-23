package entity

import "gorm.io/gorm"

type Armament struct {
	gorm.Model
	Title       string
	Qty         int
	SpaceShipID uint
}
