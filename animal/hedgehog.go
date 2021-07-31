package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	hedgehogCode    string = "hgc/hgh/hgs"
	hedgehogId      string = "hedgehog"
	hedgehogUnicode string = "🦔"
)

const (
	hedgehogFictional bool = false
	hedgehogSpecial   bool = false
)

var (
	// hedgehogNameAmericanEnglish is the name of an Hedgehog in American English.
	hedgehogNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(hedgehogId)), 0}
)

var (
	// hedgehogNames are the names of an Hedgehog in various languages.
	hedgehogNames = names{
		language.AmericanEnglish: hedgehogNameAmericanEnglish}
)

var (
	// Hedgehog is an Animal Crossing animal type.
	Hedgehog Animal = animal{
		code:      hedgehogCode,
		id:        hedgehogId,
		fictional: hedgehogFictional,
		names:     hedgehogNames,
		special:   hedgehogSpecial,
		unicode:   hedgehogUnicode}
)
