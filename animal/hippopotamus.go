package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	hippopotamusCode    string = "hip"
	hippopotamusId      string = "hippopotamus"
	hippopotamusUnicode string = "🦛"
)

const (
	hippopotamusFictional bool = false
	hippopotamusSpecial   bool = false
)

var (
	// hippopotamusNameAmericanEnglish is the name of an Hippopotamus in American English.
	hippopotamusNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(hippopotamusId)), 0}
)

var (
	// hippopotamusNames are the names of an Hippopotamus in various languages.
	hippopotamusNames = names{
		language.AmericanEnglish: hippopotamusNameAmericanEnglish}
)

var (
	// Hippopotamus is an Animal Crossing animal type.
	Hippopotamus Animal = animal{
		code:      hippopotamusCode,
		id:        hippopotamusId,
		fictional: hippopotamusFictional,
		names:     hippopotamusNames,
		special:   hippopotamusSpecial,
		unicode:   hippopotamusUnicode}
)
