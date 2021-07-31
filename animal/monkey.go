package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	monkeyCode    string = "mnk"
	monkeyId      string = "monkey"
	monkeyUnicode string = "🐒"
)

const (
	monkeyFictional bool = false
	monkeySpecial   bool = false
)

var (
	// monkey is the name of an Monkey in American English.
	monkeyNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(monkeyId)), 0}
)

var (
	// monkeyNames are the names of an Monkey in various languages.
	monkeyNames = names{
		language.AmericanEnglish: monkeyNameAmericanEnglish}
)

var (
	// Monkey is an Animal Crossing animal type.
	Monkey Animal = animal{
		code:      monkeyCode,
		id:        monkeyId,
		fictional: monkeyFictional,
		names:     monkeyNames,
		special:   monkeySpecial,
		unicode:   monkeyUnicode}
)
