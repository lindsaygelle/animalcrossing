package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/jp"
)

type Jp struct {
	jp.Jp
}

func (j Jp) Value() string {
	return "ライオン"
}
