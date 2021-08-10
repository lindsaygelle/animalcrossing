package camofrog

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg03"
)

const (
	gender string = "male"
)

const (
	id string = "camofrog"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Camofrog.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Camofrog,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
