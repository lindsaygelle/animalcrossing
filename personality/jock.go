package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	jockId string = "jock"
)

var (
	// jockAmericanEnglish is the name of Jock in American English.
	jockAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(jockId))}
)

var (
	// jockNames are the names of Jock in various languages.
	jockNames = names{
		language.AmericanEnglish: jockAmericanEnglish}
)

var (
	// Jock is an Animal Crossing personality type.
	Jock Personality = personality{
		id:    jockId,
		names: jockNames}
)
