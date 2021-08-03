package annalisa

import (
	"github.com/lindsaygelle/animalcrossing/animal/anteater"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "annalisa"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Annalisa.
	Villager = villager.Villager{
		Animal:  anteater.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
