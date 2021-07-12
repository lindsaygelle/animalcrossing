package dog

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	Dog animals.Animal = dog{}
)

type dog struct{}

func (c dog) Name() string {
	return "Dog"
}

func (c dog) Special() bool {
	return false
}

func (c dog) Translation() translations.Translation {
	return translation{}
}
