package personality

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	lazyId string = "lazy"
)

var (
	lazyAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(lazyId))}
)

var (
	lazyNames = names{
		language.AmericanEnglish: lazyAmericanEnglish}
)

var (
	// Lazy is an Animal Crossing personality type.
	Lazy Personality = personality{
		id:    lazyId,
		names: lazyNames}
)
