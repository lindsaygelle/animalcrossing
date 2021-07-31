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
	scorpiusAmericanEnglish = name{
		translation.New(language.AmericanEnglish, scorpiusId)}
)

var (
	scorpiusNames = names{
		language.AmericanEnglish: scorpiusAmericanEnglish}
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
