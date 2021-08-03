package diana

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "diana"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Diana.
	Villager = villager.Villager{
		Animal:  deer.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
