package paolo

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "elp05"
)

const (
	gender string = "male"
)

const (
	id string = "paolo"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Paolo.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Paolo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
