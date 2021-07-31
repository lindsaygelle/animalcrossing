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
	crankyAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(crankyId))}
)

var (
	crankyNames = names{
		language.AmericanEnglish: crankyAmericanEnglish}
)

var (
	// Cranky is an Animal Crossing personality type.
	Cranky Personality = personality{
		id:    crankyId,
		names: crankyNames}
)
