package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	walrusCode    string = "wls/wrl"
	walrusId      string = "walrus"
	walrusUnicode string = ""
)

const (
	walrusFictional bool = false
	walrusSpecial   bool = false
)

var (
	// walrus is the name of an Walrus in American English.
	walrusNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(walrusId)), 0}
)

var (
	// walrusNames are the names of an Walrus in various languages.
	walrusNames = names{
		language.AmericanEnglish: walrusNameAmericanEnglish}
)

var (
	// Walrus is an Animal Crossing animal type.
	Walrus Animal = animal{
		code:      walrusCode,
		id:        walrusId,
		fictional: walrusFictional,
		names:     walrusNames,
		special:   walrusSpecial,
		unicode:   walrusUnicode}
)
