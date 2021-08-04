package gladys

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "gladys"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gladys.
	Villager = villager.Villager{
		Animal:      ostrich.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gladys,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
