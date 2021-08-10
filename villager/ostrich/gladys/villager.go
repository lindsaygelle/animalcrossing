package gladys

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ost01"
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
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gladys,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
