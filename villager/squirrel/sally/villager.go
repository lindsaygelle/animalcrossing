package sally

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "sally"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sally.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Sally,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
