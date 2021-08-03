package nan

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "nan"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Nan.
	Villager = villager.Villager{
		Animal:  goat.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
