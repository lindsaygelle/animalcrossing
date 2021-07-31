package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	smugId string = "smug"
)

var (
	// smugAmericanEnglish is the name of Smug in American English.
	smugAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, strings.Title(smugId))
)

var (
	// smugNames are the names of Smug in various languages.
	smugNames = translation.NewTranslations(
		smugAmericanEnglish)
)

var (
	// Smug is an Animal Crossing personality type.
	Smug Personality = personality{
		id:    smugId,
		names: smugNames}
)
