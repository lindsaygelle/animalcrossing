package bear_cub

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	BearCub animals.Animal = bearCub{}
)

type bearCub struct{}

func (b bearCub) Name() string {
	return "Bear Cub"
}

func (b bearCub) Special() bool {
	return false
}

func (b bearCub) Translation() translations.Translation {
	return translation{}
}
