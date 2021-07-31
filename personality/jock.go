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
	jockAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(jockId))}
)

var (
	jockNames = names{
		language.AmericanEnglish: jockAmericanEnglish}
)

var (
	// Jock is an Animal Crossing personality type.
	Jock Personality = personality{
		id:    jockId,
		names: jockNames}
)
