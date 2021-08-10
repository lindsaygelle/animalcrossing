package angus

import (
	"github.com/lindsaygelle/animalcrossing/animal/bull"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "bul00"
)

const (
	gender string = "male"
)

const (
	id string = "angus"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Angus.
	Villager = villager.Villager{
		Animal:      bull.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Angus,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
