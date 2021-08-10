package leonardo

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "tig04"
)

const (
	gender string = "male"
)

const (
	id string = "leonardo"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Leonardo.
	Villager = villager.Villager{
		Animal:      tiger.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Leonardo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
