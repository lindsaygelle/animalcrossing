package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	sisterlyId string = "sisterly"
)

var (
	// sisterlyAmericanEnglish is the name of Sisterly in American English.
	sisterlyAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, strings.Title(sisterlyId))
)

var (
	// sisterlyNames are the names of Sisterly in various languages.
	sisterlyNames = translation.NewTranslations(
		sisterlyAmericanEnglish)
)

var (
	// Sisterly is an Animal Crossing personality type.
	Sisterly Personality = personality{
		id:    sisterlyId,
		names: sisterlyNames}
)
