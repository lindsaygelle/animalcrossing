package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	horseCode    string = "hrs"
	horseId      string = "horse"
	horseUnicode string = "🐎"
)

const (
	horseFictional bool = false
	horseSpecial   bool = false
)

var (
	// horseNameAmericanEnglish is the name of an Horse in American English.
	horseNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(horseId)), 0}
)

var (
	// horseNames are the names of an Horse in various languages.
	horseNames = names{
		language.AmericanEnglish: horseNameAmericanEnglish}
)

var (
	// Horse is an Animal Crossing animal type.
	Horse Animal = animal{
		code:      horseCode,
		id:        horseId,
		fictional: horseFictional,
		names:     horseNames,
		special:   horseSpecial,
		unicode:   horseUnicode}
)
