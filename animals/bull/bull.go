package bull

import (
	"github.com/lindsaygelle/animalcrossing/animals"
	"github.com/lindsaygelle/animalcrossing/translations"
)

var (
	Bull animals.Animal = bull{}
)

type bull struct{}

func (b bull) Name() string {
	return "Bull"
}

func (b bull) Special() bool {
	return false
}

func (b bull) Translation() translations.Translation {
	return translation{}
}
