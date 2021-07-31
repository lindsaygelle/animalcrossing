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
	sisterlyAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(sisterlyId))}
)

var (
	sisterlyNames = names{
		language.AmericanEnglish: sisterlyAmericanEnglish}
)

var (
	// Sisterly is an Animal Crossing personality type.
	Sisterly Personality = personality{
		id:    sisterlyId,
		names: sisterlyNames}
)
