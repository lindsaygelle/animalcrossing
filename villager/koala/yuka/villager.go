package yuka

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "kal00"
)

const (
	gender string = "female"
)

const (
	id string = "yuka"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Yuka.
	Villager = villager.Villager{
		Animal:      koala.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Yuka,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
