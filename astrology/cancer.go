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
	cancerAmericanEnglish = name{
		translation.New(language.AmericanEnglish, cancerId)}
)

var (
	cancerNames = names{
		language.AmericanEnglish: cancerAmericanEnglish}
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
