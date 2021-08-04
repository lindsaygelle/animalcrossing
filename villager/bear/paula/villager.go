package paula

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "paula"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Paula.
	Villager = villager.Villager{
		Animal:      bear.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Paula,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
