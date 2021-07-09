package deer

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	D animals.Animal = cow{}
)

type cow struct{}

func (c cow) Name() string {
	return "Cow"
}

func (c cow) Special() bool {
	return false
}

func (c cow) Translation() translations.Translation {
	return translation{}
}
