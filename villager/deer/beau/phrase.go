package beau

import "golang.org/x/text/language"

const (
	phraseAmericanEnglish      string = "saltlick"
	phraseCanadianFrench       string = "feufeuille"
	phraseDutch                string = "zoutsteen"
	phraseFrench               string = "feufeuille"
	phraseGerman               string = "knospel"
	phraseItalian              string = "gnam gnam"
	phraseJapanese             string = "おろおろ"
	phraseLatinAmericanSpanish string = "babum"
	phraseKorean               string = "우왕"
	phraseRussian              string = "соль-соль"
	phraseSpanish              string = "babum"
	phraseSimplifiedChinese    string = "怎么办"
	phraseTraditionalChinese   string = "怎麼辦"
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
