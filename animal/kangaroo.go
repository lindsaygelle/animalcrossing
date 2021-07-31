package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	kangarooCode    string = "kgr"
	kangarooId      string = "kangaroo"
	kangarooUnicode string = "🦘"
)

const (
	kangarooFictional bool = false
	kangarooSpecial   bool = false
)

var (
	// kangarooNameAmericanEnglish is the name of an Kangaroo in American English.
	kangarooNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(kangarooId)), 0}
)

var (
	// kangarooNames are the names of an Kangaroo in various languages.
	kangarooNames = names{
		language.AmericanEnglish: kangarooNameAmericanEnglish}
)

var (
	// Kangaroo is an Animal Crossing animal type.
	Kangaroo Animal = animal{
		code:      kangarooCode,
		id:        kangarooId,
		fictional: kangarooFictional,
		names:     kangarooNames,
		special:   kangarooSpecial,
		unicode:   kangarooUnicode}
)
