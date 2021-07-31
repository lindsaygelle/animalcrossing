package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	sagittariusId      string = "sagittarius"
	sagittariusSymbol  string = "♐︎"
	sagittariusUnicode string = "♐"
)

const (
	sagittariusLongitude uint16 = 240
)

var (
	// sagittariusAmericanEnglish is the name of Sagittarius in American English.
	sagittariusAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, sagittariusId)
)

var (
	// sagittariusNames are the names of Sagittarius in various languages.
	sagittariusNames = translation.NewTranslations(
		sagittariusAmericanEnglish)
)

var (
	// Sagittarius is the astrological information for the star sign Sagittarius.
	Sagittarius Astrology = astrology{
		id:         sagittariusId,
		longtitude: sagittariusLongitude,
		names:      sagittariusNames,
		symbol:     sagittariusSymbol,
		unicode:    sagittariusUnicode}
)
