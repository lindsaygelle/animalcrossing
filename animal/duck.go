package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	duckCode    string = "duk"
	duckId      string = "duck"
	duckUnicode string = "🦆"
)

const (
	duckFictional bool = false
	duckSpecial   bool = false
)

var (
	// duckNameAmericanEnglish is the name of an Duck in American English.
	duckNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(duckId)), 0}
)

var (
	// duckNames are the names of an Duck in various languages.
	duckNames = names{
		language.AmericanEnglish: duckNameAmericanEnglish}
)

var (
	// Duck is an Animal Crossing animal type.
	Duck Animal = animal{
		code:      duckCode,
		id:        duckId,
		fictional: duckFictional,
		names:     duckNames,
		special:   duckSpecial,
		unicode:   duckUnicode}
)
