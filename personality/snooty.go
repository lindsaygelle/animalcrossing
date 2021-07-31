package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	snootyId string = "snooty"
)

var (
	// snootyAmericanEnglish is the name of Snooty in American English.
	snootyAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(snootyId))}
)

var (
	// smugNames are the names of Snooty in various languages.
	snootyNames = names{
		language.AmericanEnglish: snootyAmericanEnglish}
)

var (
	// Snooty is an Animal Crossing personality type.
	Snooty Personality = personality{
		id:    snootyId,
		names: snootyNames}
)
