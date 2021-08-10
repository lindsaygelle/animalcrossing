package bertha

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hip03"
)

const (
	gender string = "female"
)

const (
	id string = "bertha"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bertha.
	Villager = villager.Villager{
		Animal:      hippopotamus.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bertha,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
