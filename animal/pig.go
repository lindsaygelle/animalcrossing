package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	pigCode    string = "pig"
	pigId      string = "pig"
	pigUnicode string = "🐖"
)

const (
	pigFictional bool = false
	pigSpecial   bool = false
)

var (
	// pigNameAmericanEnglish is the name of an Pig in American English.
	pigNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(pigId)), 0}
)

var (
	// pigNames are the names of an Pig in various languages.
	pigNames = names{
		language.AmericanEnglish: pigNameAmericanEnglish}
)

var (
	// Pig is an Animal Crossing animal type.
	Pig Animal = animal{
		code:      pigCode,
		id:        pigId,
		fictional: pigFictional,
		names:     pigNames,
		special:   pigSpecial,
		unicode:   pigUnicode}
)
