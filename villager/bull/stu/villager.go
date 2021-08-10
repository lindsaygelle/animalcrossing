package stu

import (
	"github.com/lindsaygelle/animalcrossing/animal/bull"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "bul03"
)

const (
	gender string = "male"
)

const (
	id string = "stu"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Stu.
	Villager = villager.Villager{
		Animal:      bull.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Stu,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
