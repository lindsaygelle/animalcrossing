package chief

import (
	"github.com/lindsaygelle/animalcrossing/animal/wolf"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "wol00"
)

const (
	gender string = "male"
)

const (
	id string = "chief"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Chief.
	Villager = villager.Villager{
		Animal:      wolf.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Chief,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
