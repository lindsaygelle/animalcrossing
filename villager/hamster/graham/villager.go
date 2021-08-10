package graham

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ham02"
)

const (
	gender string = "male"
)

const (
	id string = "graham"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Graham.
	Villager = villager.Villager{
		Animal:      hamster.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Graham,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
