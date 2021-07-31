package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	bearcubCode    string = "cbr"
	bearcubId      string = "bearcub"
	bearcubUnicode string = "🐮"
)

const (
	bearcubFictional bool = false
	bearcubSpecial   bool = false
)

var (
	// bearcubNameAmericanEnglish is the name of an Bearcub in American English.
	bearcubNameAmericanEnglish = name{
		translation.NewTranslation(language.AmericanEnglish, strings.Title(bearcubId)), 0}
)

var (
	// bearcubNames are the names of an Bearcub in various languages.
	bearcubNames = names{
		language.AmericanEnglish: bearcubNameAmericanEnglish}
)

var (
	// Bearcub is an Animal Crossing animal type.
	Bearcub Animal = animal{
		code:      bearcubCode,
		id:        bearcubId,
		fictional: bearcubFictional,
		names:     bearcubNames,
		special:   bearcubSpecial,
		unicode:   bearcubUnicode}
)
