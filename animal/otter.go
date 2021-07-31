package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	otterCode    string = "otg/ott/seo"
	otterId      string = "otter"
	otterUnicode string = "🦦"
)

const (
	otterFictional bool = false
	otterSpecial   bool = false
)

var (
	// otterNameAmericanEnglish is the name of an Otter in American English.
	otterNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(otterId)), 0}
)

var (
	// otterNames are the names of an Otter in various languages.
	otterNames = names{
		language.AmericanEnglish: otterNameAmericanEnglish}
)

var (
	// Otter is an Animal Crossing animal type.
	Otter Animal = animal{
		code:      otterCode,
		id:        otterId,
		fictional: otterFictional,
		names:     otterNames,
		special:   otterSpecial,
		unicode:   otterUnicode}
)
