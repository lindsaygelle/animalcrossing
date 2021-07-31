package villager

import "github.com/lindsaygelle/animalcrossing/translation"

// Name is an Animal Crossing villager name.
type Name interface {
	translation.Translation
}

// name implements Name.
type name struct {
	translation.Translation
}

var (
	_ Name = name{}
)
