package alligator

import "github.com/lindsaygelle/animalcrossing/languages"

type Alligator struct{}

func (a Alligator) Id() string {
	return "alligator"
}

func (a Alligator) Special() bool {
	return false
}

func (a Alligator) Translations() []languages.Translation {
	return []languages.Translation{de{}, en{}, es{}, fr{}, it{}, jp{}, ko{}, nl{}, ru{}, zh{}}
}
