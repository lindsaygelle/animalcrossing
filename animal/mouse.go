package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	mouseCode    string = "mus"
	mouseId      string = "mouse"
	mouseUnicode string = "🐁"
)

const (
	mouseFictional bool = false
	mouseSpecial   bool = false
)

var (
	// mouseNameAmericanEnglish is the name of an Mouse in American English.
	mouseNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(mouseId)), 0}
)

var (
	// mouseNames are the names of an Mouse in various languages.
	mouseNames = names{
		language.AmericanEnglish: mouseNameAmericanEnglish}
)

var (
	// Mouse is an Animal Crossing animal type.
	Mouse Animal = animal{
		code:      mouseCode,
		id:        mouseId,
		fictional: mouseFictional,
		names:     mouseNames,
		special:   mouseSpecial,
		unicode:   mouseUnicode}
)
