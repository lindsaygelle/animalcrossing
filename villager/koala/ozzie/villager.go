package ozzie

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "ozzie"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ozzie.
	Villager = villager.Villager{
		Animal:  koala.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
