package apollo

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "apollo"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Apollo.
	Villager = villager.Villager{
		Animal:      eagle.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Apollo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
