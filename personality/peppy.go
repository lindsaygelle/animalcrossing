package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	peppyId string = "peppy"
)

var (
	peppyAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(peppyId))}
)

var (
	peppyNames = names{
		language.AmericanEnglish: peppyAmericanEnglish}
)

var (
	// Peppy is an Animal Crossing personality type.
	Peppy Personality = personality{
		id:    peppyId,
		names: peppyNames}
)
