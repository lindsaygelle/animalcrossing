package alpaca

import "github.com/lindsaygelle/animalcrossing/languages"

type Alpaca struct{}

func (a Alpaca) Id() string {
	return "alpaca"
}

func (a Alpaca) Special() bool {
	return true
}

func (a Alpaca) Translations() []languages.Translation {
	return []languages.Translation{}
}
