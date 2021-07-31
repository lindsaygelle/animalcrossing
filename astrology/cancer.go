package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	cancerId      string = "cancer"
	cancerSymbol  string = "♋︎"
	cancerUnicode string = "♋"
)

const (
	cancerLongitude uint16 = 90
)

var (
	// cancerAmericanEnglish is the name of Cancer in American English.
	cancerAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, cancerId)
)

var (
	// cancerNames are the names of Cancer in various languages.
	cancerNames = translation.NewTranslations(
		cancerAmericanEnglish)
)

var (
	// Cancer is the astrological information for the star sign Cancer.
	Cancer Astrology = astrology{
		id:         cancerId,
		longtitude: cancerLongitude,
		names:      cancerNames,
		symbol:     cancerSymbol,
		unicode:    cancerUnicode}
)
