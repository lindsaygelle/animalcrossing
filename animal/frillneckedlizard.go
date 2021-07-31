package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	frillneckedlizardCode    string = "liz"
	frillneckedlizardId      string = "frillneckedlizard"
	frillneckedlizardUnicode string = "🦎"
)

const (
	frillneckedlizardFictional bool = false
	frillneckedlizardSpecial   bool = false
)

var (
	// frillneckedlizardNameAmericanEnglish is the name of an FrillneckedLizard in American English.
	frillneckedlizardNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(frillneckedlizardId)), 0}
)

var (
	// frillneckedlizardNames are the names of an FrillneckedLizard in various languages.
	frillneckedlizardNames = names{
		language.AmericanEnglish: frillneckedlizardNameAmericanEnglish}
)

var (
	// FrillneckedLizard is an Animal Crossing animal type.
	FrillneckedLizard Animal = animal{
		code:      frillneckedlizardCode,
		id:        frillneckedlizardId,
		fictional: frillneckedlizardFictional,
		names:     frillneckedlizardNames,
		special:   frillneckedlizardSpecial,
		unicode:   frillneckedlizardUnicode}
)
