package ozzie

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "kal05"
)

const (
	gender string = "male"
)

const (
	id string = "ozzie"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ozzie.
	Villager = villager.Villager{
		Animal:      koala.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ozzie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
