package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	rabbitCode    string = "rbt"
	rabbitId      string = "rabbit"
	rabbitUnicode string = "🐇"
)

const (
	rabbitFictional bool = false
	rabbitSpecial   bool = false
)

var (
	// rabbit is the name of an Rabbit in American English.
	rabbitNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(rabbitId)), 0}
)

var (
	// rabbitNames are the names of an Rabbit in various languages.
	rabbitNames = names{
		language.AmericanEnglish: rabbitNameAmericanEnglish}
)

var (
	// Rabbit is an Animal Crossing animal type.
	Rabbit Animal = animal{
		code:      rabbitCode,
		id:        rabbitId,
		fictional: rabbitFictional,
		names:     rabbitNames,
		special:   rabbitSpecial,
		unicode:   rabbitUnicode}
)
