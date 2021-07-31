package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	aquariusId      string = "aquarius"
	aquariusSymbol  string = "♒︎"
	aquariusUnicode string = "♒"
)

const (
	aquariusLongitude uint16 = 300
)

var (
	// aquariusAmericanEnglishName is the name of Aquarius in American English.
	aquariusAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, aquariusId)
)

var (
	// aquariusNames are the names of Aquarius in various languages.
	aquariusNames = translation.NewTranslations(
		aquariusAmericanEnglish)
)

var (
	// Aquarius is the astrological information for the star sign Aquarius.
	Aquarius Astrology = astrology{
		id:         aquariusId,
		longtitude: aquariusLongitude,
		names:      aquariusNames,
		symbol:     aquariusSymbol,
		unicode:    aquariusUnicode}
)
