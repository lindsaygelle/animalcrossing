package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	turkeyCode    string = "tuk"
	turkeyId      string = "turkey"
	turkeyUnicode string = "🦃"
)

const (
	turkeyFictional bool = false
	turkeySpecial   bool = false
)

var (
	// turkeyNameAmericanEnglish is the name of an Turkey in American English.
	turkeyNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(turkeyId)), 0}
)

var (
	// turkeyNames are the names of an Turkey in various languages.
	turkeyNames = names{
		language.AmericanEnglish: turkeyNameAmericanEnglish}
)

var (
	// Turkey is an Animal Crossing animal type.
	Turkey Animal = animal{
		code:      turkeyCode,
		id:        turkeyId,
		fictional: turkeyFictional,
		names:     turkeyNames,
		special:   turkeySpecial,
		unicode:   turkeyUnicode}
)
