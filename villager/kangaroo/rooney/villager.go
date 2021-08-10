package rooney

import (
	"github.com/lindsaygelle/animalcrossing/animal/kangaroo"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "kgr09"
)

const (
	gender string = "male"
)

const (
	id string = "rooney"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rooney.
	Villager = villager.Villager{
		Animal:      kangaroo.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rooney,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
