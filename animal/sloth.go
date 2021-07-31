package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	slothCode    string = "slo"
	slothId      string = "sloth"
	slothUnicode string = "🦥"
)

const (
	slothFictional bool = false
	slothSpecial   bool = false
)

var (
	// slothNameAmericanEnglish is the name of an Sloth in American English.
	slothNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(slothId)), 0}
)

var (
	// slothNames are the names of an Sloth in various languages.
	slothNames = names{
		language.AmericanEnglish: slothNameAmericanEnglish}
)

var (
	// Sloth is an Animal Crossing animal type.
	Sloth Animal = animal{
		code:      slothCode,
		id:        slothId,
		fictional: slothFictional,
		names:     slothNames,
		special:   slothSpecial,
		unicode:   slothUnicode}
)
