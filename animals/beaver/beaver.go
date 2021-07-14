package beaver

import "github.com/lindsaygelle/animalcrossing/languages"

type Beaver struct{}

func (b Beaver) Id() string {
	return "beaver"
}

func (b Beaver) Special() bool {
	return true
}

func (b Beaver) Translations() []languages.Translation {
	return []languages.Translation{}
}
