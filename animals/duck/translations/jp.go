package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/jp"
)

type Jp struct {
	jp.Jp
	genders.None
}

func (j Jp) Value() string {
	return "アヒル"
}
