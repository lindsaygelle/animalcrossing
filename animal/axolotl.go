package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	axolotlCode    string = "upa"
	axolotlId      string = "axolotl"
	axolotlUnicode string = ""
)

const (
	axolotlFictional bool = false
	axolotlSpecial   bool = false
)

var (
	// axolotl is the name of an Axolotl in American English.
	axolotlNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(axolotlId)), 0}
)

var (
	// axolotlNames are the names of an Axolotl in various languages.
	axolotlNames = names{
		language.AmericanEnglish: axolotlNameAmericanEnglish}
)

var (
	// Axolotl is an Animal Crossing animal type.
	Axolotl Animal = animal{
		code:      axolotlCode,
		id:        axolotlId,
		fictional: axolotlFictional,
		names:     axolotlNames,
		special:   axolotlSpecial,
		unicode:   axolotlUnicode}
)
