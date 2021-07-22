package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/jp"
)

type Jp struct {
	jp.Jp
}

func (j Jp) Id() string {
	return "none"
}

func (j Jp) Value() string {
	return "ワニ"
}
