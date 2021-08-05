package alfonso

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "crd00"
)

const (
	gender string = "male"
)

const (
	id string = "alfonso"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Alfonso.
	Villager = villager.Villager{
		Animal:      alligator.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.Alfonso,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
