package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	normalId string = "normal"
)

var (
	// normalAmericanEnglish is the name of Normal in American English.
	normalAmericanEnglish = translation.NewTranslation(
		language.AmericanEnglish, strings.Title(normalId))
)

var (
	// normalNames are the names of Normal in various languages.
	normalNames = translation.NewTranslations(
		normalAmericanEnglish)
)

var (
	// Normal is an Animal Crossing personality type.
	Normal Personality = personality{
		id:    normalId,
		names: normalNames}
)
