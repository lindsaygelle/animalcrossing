package groucho

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "grumble"
	phraseCanadianFrench       string = "grou-gnon"
	phraseDutch                string = "brom"
	phraseFrench               string = "grou-gnon"
	phraseGerman               string = "brummel"
	phraseItalian              string = "grum grum"
	phraseJapanese             string = "ワォー"
	phraseLatinAmericanSpanish string = "jorrr"
	phraseKorean               string = "어쩌라구"
	phraseRussian              string = "бур-бур"
	phraseSpanish              string = "jorrr"
	phraseSimplifiedChinese    string = "哇哇"
	phraseTraditionalChinese   string = "哇哇"
)

var (
	phrase = map[language.Tag]string{
		language.AmericanEnglish:      phraseAmericanEnglish,
		language.CanadianFrench:       phraseCanadianFrench,
		language.Dutch:                phraseDutch,
		language.French:               phraseFrench,
		language.German:               phraseGerman,
		language.Italian:              phraseItalian,
		language.Japanese:             phraseJapanese,
		language.Korean:               phraseKorean,
		language.LatinAmericanSpanish: phraseLatinAmericanSpanish,
		language.Russian:              phraseRussian,
		language.Spanish:              phraseSpanish,
		language.SimplifiedChinese:    phraseSimplifiedChinese,
		language.TraditionalChinese:   phraseTraditionalChinese}
)
