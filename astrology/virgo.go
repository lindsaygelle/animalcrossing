package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	virgoId      string = "virgo"
	virgoSymbol  string = "♍︎"
	virgoUnicode string = "♍"
)

const (
	virgoLongitude uint16 = 150
)

var (
	// virgoAmericanEnglish is the name of Virgo in American English.
	virgoAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, virgoId)
)

var (
	// virgoNames are the names of Virgo in various languages.
	virgoNames = translation.NewTranslations(
		virgoAmericanEnglish)
)

var (
	// Virgo is the astrological information for the star sign Virgo.
	Virgo Astrology = astrology{
		id:         virgoId,
		longtitude: virgoLongitude,
		names:      virgoNames,
		symbol:     virgoSymbol,
		unicode:    virgoUnicode}
)
