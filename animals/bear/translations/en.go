package translations

import (
	e "github.com/lindsaygelle/animalcrossing/languages/en"
)

type En struct {
	e.En
}

func (e En) Id() string {
	return "none"
}

func (e En) Value() string {
	return "Bear"
}
