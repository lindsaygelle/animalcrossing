package translations

import (
	"github.com/lindsaygelle/animalcrossing/genders"
	"github.com/lindsaygelle/animalcrossing/languages/ko"
)

type Ko struct {
	ko.Ko
	genders.None
}

func (k Ko) Value() string {
	return "개미핥기"
}
