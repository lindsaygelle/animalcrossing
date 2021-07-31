package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	snowCode    string = ""
	snowId      string = "snow"
	snowUnicode string = "☃️"
)

const (
	snowFictional bool = false
	snowSpecial   bool = false
)

var (
	// snowNameAmericanEnglish is the name of an Snow in American English.
	snowNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(snowId)), 0}
)

var (
	// snowNames are the names of an Snow in various languages.
	snowNames = names{
		language.AmericanEnglish: snowNameAmericanEnglish}
)

var (
	// Snow is an Animal Crossing animal type.
	Snow Animal = animal{
		code:      snowCode,
		id:        snowId,
		fictional: snowFictional,
		names:     snowNames,
		special:   snowSpecial,
		unicode:   snowUnicode}
)
