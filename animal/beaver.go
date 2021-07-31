package animal

import (
	"strings"

	"github.com/lindsaygelle/animalcrossing/translation"
	"golang.org/x/text/language"
)

const (
	beaverCode    string = "bev/bey"
	beaverId      string = "beaver"
	beaverUnicode string = "🦫"
)

const (
	beaverFictional bool = false
	beaverSpecial   bool = false
)

var (
	// beaver is the name of an Beaver in American English.
	beaverNameAmericanEnglish = name{
		translation.New(language.AmericanEnglish, strings.Title(beaverId)), 0}
)

var (
	// beaverNames are the names of an Beaver in various languages.
	beaverNames = names{
		language.AmericanEnglish: beaverNameAmericanEnglish}
)

var (
	// Beaver is an Animal Crossing animal type.
	Beaver Animal = animal{
		code:      beaverCode,
		id:        beaverId,
		fictional: beaverFictional,
		names:     beaverNames,
		special:   beaverSpecial,
		unicode:   beaverUnicode}
)
