package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	squirrelCode    string = "squ"
	squirrelId      string = "squirrel"
	squirrelUnicode string = ""
)

const (
	squirrelFictional bool = false
	squirrelSpecial   bool = false
)

var (
	// squirrelNameAmericanEnglish is the name of an Squirrel in American English.
	squirrelNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(squirrelId)), 0}
)

var (
	// squirrelNames are the names of an Squirrel in various languages.
	squirrelNames = names{
		language.AmericanEnglish: squirrelNameAmericanEnglish}
)

var (
	// Squirrel is an Animal Crossing animal type.
	Squirrel Animal = animal{
		code:      squirrelCode,
		id:        squirrelId,
		fictional: squirrelFictional,
		names:     squirrelNames,
		special:   squirrelSpecial,
		unicode:   squirrelUnicode}
)
