package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	foxCode    string = "fox"
	foxId      string = "fox"
	foxUnicode string = "🦊"
)

const (
	foxFictional bool = false
	foxSpecial   bool = false
)

var (
	// foxNameAmericanEnglish is the name of an Fox in American English.
	foxNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(foxId)), 0}
)

var (
	// foxNames are the names of an Fox in various languages.
	foxNames = names{
		language.AmericanEnglish: foxNameAmericanEnglish}
)

var (
	// Fox is an Animal Crossing animal type.
	Fox Animal = animal{
		code:      foxCode,
		id:        foxId,
		fictional: foxFictional,
		names:     foxNames,
		special:   foxSpecial,
		unicode:   foxUnicode}
)
