package benedict

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "chn01"
)

const (
	gender string = "male"
)

const (
	id string = "benedict"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Benedict.
	Villager = villager.Villager{
		Animal:      chicken.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Benedict,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
