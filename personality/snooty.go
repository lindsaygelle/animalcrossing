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
	snootyAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(snootyId))}
)

var (
	snootyNames = names{
		language.AmericanEnglish: snootyAmericanEnglish}
)

var (
	// Snooty is an Animal Crossing personality type.
	Snooty Personality = personality{
		id:    snootyId,
		names: snootyNames}
)
