package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/zh"
)

type Zh struct {
	zh.Zh
	genders.None
}

func (z Zh) Value() string {
	return "鳄鱼"
}
