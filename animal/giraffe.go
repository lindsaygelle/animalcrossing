package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	giraffeCode    string = "grf"
	giraffeId      string = "giraffe"
	giraffeUnicode string = "🦒"
)

const (
	giraffeFictional bool = false
	giraffeSpecial   bool = false
)

var (
	// giraffe is the name of an Giraffe in American English.
	giraffeNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(giraffeId)), 0}
)

var (
	// giraffeNames are the names of an Giraffe in various languages.
	giraffeNames = names{
		language.AmericanEnglish: giraffeNameAmericanEnglish}
)

var (
	// Giraffe is an Animal Crossing animal type.
	Giraffe Animal = animal{
		code:      giraffeCode,
		id:        giraffeId,
		fictional: giraffeFictional,
		names:     giraffeNames,
		special:   giraffeSpecial,
		unicode:   giraffeUnicode}
)
