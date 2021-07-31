package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	smugId string = "smug"
)

var (
	smugAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(smugId))}
)

var (
	smugNames = names{
		language.AmericanEnglish: smugAmericanEnglish}
)

var (
	// Smug is an Animal Crossing personality type.
	Smug Personality = personality{
		id:    smugId,
		names: smugNames}
)
