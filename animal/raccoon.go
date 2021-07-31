package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	raccoonCode    string = "rcn/rco/lrn/rct"
	raccoonId      string = "raccoon"
	raccoonUnicode string = "🦝"
)

const (
	raccoonFictional bool = false
	raccoonSpecial   bool = false
)

var (
	// raccoonNameAmericanEnglish is the name of an Raccoon in American English.
	raccoonNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(raccoonId)), 0}
)

var (
	// raccoonNames are the names of an Raccoon in various languages.
	raccoonNames = names{
		language.AmericanEnglish: raccoonNameAmericanEnglish}
)

var (
	// Raccoon is an Animal Crossing animal type.
	Raccoon Animal = animal{
		code:      raccoonCode,
		id:        raccoonId,
		fictional: raccoonFictional,
		names:     raccoonNames,
		special:   raccoonSpecial,
		unicode:   raccoonUnicode}
)
