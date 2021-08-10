package avery

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pbr05"
)

const (
	gender string = "male"
)

const (
	id string = "avery"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Avery.
	Villager = villager.Villager{
		Animal:      eagle.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Avery,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
