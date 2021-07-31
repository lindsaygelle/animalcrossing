package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	reindeerCode    string = "snt/rei"
	reindeerId      string = "reindeer"
	reindeerUnicode string = "🦌"
)

const (
	reindeerFictional bool = false
	reindeerSpecial   bool = false
)

var (
	// reindeerNameAmericanEnglish is the name of an Reindeer in American English.
	reindeerNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(reindeerId)), 0}
)

var (
	// reindeerNames are the names of an Reindeer in various languages.
	reindeerNames = names{
		language.AmericanEnglish: reindeerNameAmericanEnglish}
)

var (
	// Reindeer is an Animal Crossing animal type.
	Reindeer Animal = animal{
		code:      reindeerCode,
		id:        reindeerId,
		fictional: reindeerFictional,
		names:     reindeerNames,
		special:   reindeerSpecial,
		unicode:   reindeerUnicode}
)
