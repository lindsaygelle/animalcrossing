package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	wolfCode    string = "wol"
	wolfId      string = "wolf"
	wolfUnicode string = "🐺"
)

const (
	wolfFictional bool = false
	wolfSpecial   bool = false
)

var (
	// wolfNameAmericanEnglish is the name of an Wolf in American English.
	wolfNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(wolfId)), 0}
)

var (
	// wolfNames are the names of an Wolf in various languages.
	wolfNames = names{
		language.AmericanEnglish: wolfNameAmericanEnglish}
)

var (
	// Wolf is an Animal Crossing animal type.
	Wolf Animal = animal{
		code:      wolfCode,
		id:        wolfId,
		fictional: wolfFictional,
		names:     wolfNames,
		special:   wolfSpecial,
		unicode:   wolfUnicode}
)
