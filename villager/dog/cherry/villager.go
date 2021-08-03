package cherry

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "cherry"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cherry.
	Villager = villager.Villager{
		Animal:  dog.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
