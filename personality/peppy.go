package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	peppyId string = "peppy"
)

var (
	// peppyAmericanEnglish is the name of Peppy in American English.
	peppyAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, strings.Title(peppyId))
)

var (
	// peppyNames are the names of Peppy in various languages.
	peppyNames = translation.NewTranslations(
		peppyAmericanEnglish)
)

var (
	// Peppy is an Animal Crossing personality type.
	Peppy Personality = personality{
		id:    peppyId,
		names: peppyNames}
)
