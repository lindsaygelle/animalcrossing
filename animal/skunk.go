package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	skunkCode    string = "skk"
	skunkId      string = "skunk"
	skunkUnicode string = "🦨"
)

const (
	skunkFictional bool = false
	skunkSpecial   bool = false
)

var (
	// skunk is the name of an Skunk in American English.
	skunkNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(skunkId)), 0}
)

var (
	// skunkNames are the names of an Skunk in various languages.
	skunkNames = names{
		language.AmericanEnglish: skunkNameAmericanEnglish}
)

var (
	// Skunk is an Animal Crossing animal type.
	Skunk Animal = animal{
		code:      skunkCode,
		id:        skunkId,
		fictional: skunkFictional,
		names:     skunkNames,
		special:   skunkSpecial,
		unicode:   skunkUnicode}
)
