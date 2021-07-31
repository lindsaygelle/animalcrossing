package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	tigerCode    string = "tig"
	tigerId      string = "tiger"
	tigerUnicode string = "🐅"
)

const (
	tigerFictional bool = false
	tigerSpecial   bool = false
)

var (
	// tigerNameAmericanEnglish is the name of an Tiger in American English.
	tigerNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(tigerId)), 0}
)

var (
	// tigerNames are the names of an Tiger in various languages.
	tigerNames = names{
		language.AmericanEnglish: tigerNameAmericanEnglish}
)

var (
	// Tiger is an Animal Crossing animal type.
	Tiger Animal = animal{
		code:      tigerCode,
		id:        tigerId,
		fictional: tigerFictional,
		names:     tigerNames,
		special:   tigerSpecial,
		unicode:   tigerUnicode}
)
