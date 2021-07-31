package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	chickenCode    string = "chn"
	chickenId      string = "chicken"
	chickenUnicode string = "🐔"
)

const (
	chickenFictional bool = false
	chickenSpecial   bool = false
)

var (
	// chickenNameAmericanEnglish is the name of an Chicken in American English.
	chickenNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(chickenId)), 0}
)

var (
	// chickenNames are the names of an Chicken in various languages.
	chickenNames = names{
		language.AmericanEnglish: chickenNameAmericanEnglish}
)

var (
	// Chicken is an Animal Crossing animal type.
	Chicken Animal = animal{
		code:      chickenCode,
		id:        chickenId,
		fictional: chickenFictional,
		names:     chickenNames,
		special:   chickenSpecial,
		unicode:   chickenUnicode}
)
