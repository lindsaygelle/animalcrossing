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
	normalAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(normalId))}
)

var (
	normalNames = names{
		language.AmericanEnglish: normalAmericanEnglish}
)

var (
	// Normal is an Animal Crossing personality type.
	Normal Personality = personality{
		id:    normalId,
		names: normalNames}
)
