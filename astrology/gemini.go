package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	geminiId      string = "gemini"
	geminiSymbol  string = "♊︎"
	geminiUnicode string = "♊"
)

const (
	geminiLongitude uint16 = 60
)

var (
	// geminiAmericanEnglish is the name of Gemini in American English.
	geminiAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, geminiId)
)

var (
	// geminiNames are the names of Gemini in various languages.
	geminiNames = translation.NewTranslations(
		geminiAmericanEnglish)
)

var (
	// Gemini is the astrological information for the star sign Gemini.
	Gemini Astrology = astrology{
		id:         geminiId,
		longtitude: geminiLongitude,
		names:      geminiNames,
		symbol:     geminiSymbol,
		unicode:    geminiUnicode}
)
