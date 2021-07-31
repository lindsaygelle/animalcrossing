package animal

import "github.com/lindsaygelle/animalcrossing/translation"

// Name is an Animal Crossing animal name.
type Name interface {
	translation.Translation

	Gender() uint8
}

// names implements Names.
type name struct {
	translation.Translation

	gender uint8
}

func (v name) Gender() uint8 {
	return v.gender
}

var (
	_ Name = name{}
)
