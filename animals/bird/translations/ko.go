package translations

import (
	"github.com/lindsaygelle/animalcrossing/languages/ko"
)

type Ko struct {
	ko.Ko
}

func (k Ko) Id() string {
	return "none"
}

func (k Ko) Value() string {
	return "새"
}
