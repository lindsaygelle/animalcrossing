package croque

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg17"
)

const (
	gender string = "male"
)

const (
	id string = "croque"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Croque.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Croque,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
