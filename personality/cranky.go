package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	crankyId string = "cranky"
)

var (
	// crankyAmericanEnglish is the name of Cranky in American English.
	crankyAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(crankyId))}
)

var (
	// crankyNames are the names of Cranky in various languages.
	crankyNames = names{
		language.AmericanEnglish: crankyAmericanEnglish}
)

var (
	// Cranky is an Animal Crossing personality type.
	Cranky Personality = personality{
		id:    crankyId,
		names: crankyNames}
)
