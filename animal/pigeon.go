package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	pigeonCode    string = "pge"
	pigeonId      string = "pigeon"
	pigeonUnicode string = ""
)

const (
	pigeonFictional bool = false
	pigeonSpecial   bool = false
)

var (
	// pigeon is the name of an Pigeon in American English.
	pigeonNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(pigeonId)), 0}
)

var (
	// pigeonNames are the names of an Pigeon in various languages.
	pigeonNames = names{
		language.AmericanEnglish: pigeonNameAmericanEnglish}
)

var (
	// Pigeon is an Animal Crossing animal type.
	Pigeon Animal = animal{
		code:      pigeonCode,
		id:        pigeonId,
		fictional: pigeonFictional,
		names:     pigeonNames,
		special:   pigeonSpecial,
		unicode:   pigeonUnicode}
)
