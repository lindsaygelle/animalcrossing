package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	bearCode    string = "bea"
	bearId      string = "bear"
	bearUnicode string = "🐻"
)

const (
	bearFictional bool = false
	bearSpecial   bool = false
)

var (
	// bearNameAmericanEnglish is the name of an Bear in American English.
	bearNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(bearId)), 0}
)

var (
	// bearNames are the names of an Bear in various languages.
	bearNames = names{
		language.AmericanEnglish: bearNameAmericanEnglish}
)

var (
	// Bear is an Animal Crossing animal type.
	Bear Animal = animal{
		code:      bearCode,
		id:        bearId,
		fictional: bearFictional,
		names:     bearNames,
		special:   bearSpecial,
		unicode:   bearUnicode}
)
