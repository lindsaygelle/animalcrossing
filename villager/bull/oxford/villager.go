package oxford

import (
	"github.com/lindsaygelle/animalcrossing/animal/bull"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "oxford"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Oxford.
	Villager = villager.Villager{
		Animal:      bull.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Oxford,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
