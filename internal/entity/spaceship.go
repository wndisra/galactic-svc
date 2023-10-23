package entity

import "gorm.io/gorm"

type SpaceShip struct {
	gorm.Model
	Name      string
	Class     string
	Armaments []Armament
	Crew      int64
	Image     string
	Value     float64
	Status    string
}
