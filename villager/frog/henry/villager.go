package henry

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "henry"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Henry.
	Villager = villager.Villager{
		Animal:  frog.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
