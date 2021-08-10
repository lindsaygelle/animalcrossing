package marcie

import (
	"github.com/lindsaygelle/animalcrossing/animal/kangaroo"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "kgr10"
)

const (
	gender string = "female"
)

const (
	id string = "marcie"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Marcie.
	Villager = villager.Villager{
		Animal:      kangaroo.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Marcie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
