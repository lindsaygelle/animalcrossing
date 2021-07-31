package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	sheepCode    string = "shp"
	sheepId      string = "sheep"
	sheepUnicode string = "🐑"
)

const (
	sheepFictional bool = false
	sheepSpecial   bool = false
)

var (
	// sheepNameAmericanEnglish is the name of an Sheep in American English.
	sheepNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(sheepId)), 0}
)

var (
	// sheepNames are the names of an Sheep in various languages.
	sheepNames = names{
		language.AmericanEnglish: sheepNameAmericanEnglish}
)

var (
	// Sheep is an Animal Crossing animal type.
	Sheep Animal = animal{
		code:      sheepCode,
		id:        sheepId,
		fictional: sheepFictional,
		names:     sheepNames,
		special:   sheepSpecial,
		unicode:   sheepUnicode}
)
