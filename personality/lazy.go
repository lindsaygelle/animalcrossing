package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	lazyId string = "lazy"
)

var (
	// lazyAmericanEnglish is the name of Lazy in American English.
	lazyAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, strings.Title(lazyId))
)

var (
	// lazyNames are the names of Lazy in various languages.
	lazyNames = translation.NewTranslations(
		lazyAmericanEnglish)
)

var (
	// Lazy is an Animal Crossing personality type.
	Lazy Personality = personality{
		id:    lazyId,
		names: lazyNames}
)
