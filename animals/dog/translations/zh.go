package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/zh"
)

type Zh struct {
	zh.Zh
}

func (z Zh) Id() string {
	return "none"
}

func (z Zh) Value() string {
	return "狗"
}
