package astrology

import (
	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	scorpiusId      string = "scorpius"
	scorpiusSymbol  string = "♏︎"
	scorpiusUnicode string = "♏"
)

const (
	scorpiusLongitude uint16 = 210
)

var (
	// scorpiusAmericanEnglish is the name of Scorpius in American English.
	scorpiusAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, scorpiusId)
)

var (
	// scorpiusNames are the names of Scorpius in various languages.
	scorpiusNames = translation.NewTranslations(
		scorpiusAmericanEnglish)
)

var (
	// Scorpius is the astrological information for the star sign Scorpius.
	Scorpius Astrology = astrology{
		id:         scorpiusId,
		longtitude: scorpiusLongitude,
		names:      scorpiusNames,
		symbol:     scorpiusSymbol,
		unicode:    scorpiusUnicode}
)
