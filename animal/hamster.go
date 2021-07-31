package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	hamsterCode    string = "ham"
	hamsterId      string = "hamster"
	hamsterUnicode string = "🐹"
)

const (
	hamsterFictional bool = false
	hamsterSpecial   bool = false
)

var (
	// hamster is the name of an Hamster in American English.
	hamsterNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(hamsterId)), 0}
)

var (
	// hamsterNames are the names of an Hamster in various languages.
	hamsterNames = names{
		language.AmericanEnglish: hamsterNameAmericanEnglish}
)

var (
	// Hamster is an Animal Crossing animal type.
	Hamster Animal = animal{
		code:      hamsterCode,
		id:        hamsterId,
		fictional: hamsterFictional,
		names:     hamsterNames,
		special:   hamsterSpecial,
		unicode:   hamsterUnicode}
)
