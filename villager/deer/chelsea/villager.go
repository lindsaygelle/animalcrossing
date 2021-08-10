package chelsea

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "der10"
)

const (
	gender string = "female"
)

const (
	id string = "chelsea"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Chelsea.
	Villager = villager.Villager{
		Animal:      deer.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Chelsea,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
